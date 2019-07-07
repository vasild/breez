package chainservice

import (
	"sync"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/breez/breez/config"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcwallet/walletdb"
	"github.com/lightninglabs/neutrino/headerfs"
)

const (
	blockHeaderLength  = 80
	filterHeaderLength = 32
)

var (
	bootstrapMu	sync.Mutex
)

func NeedsBootstrap(workingDir string) (bool, error) {
	bootstrapMu.Lock()
	defer bootstrapMu.Unlock()

	config, err := config.GetConfig(workingDir)
	if err != nil {
		return false, err
	}

	params, err := chainParams(config.Network)
	if err != nil {
		return false, err
	}

	//create temporary neturino db.
	neutrinoDataDir := neutrinoDataDir(workingDir, config.Network)
	neutrinoDB := path.Join(neutrinoDataDir, "neutrino.db")	
	if err := os.MkdirAll(neutrinoDataDir, 0700); err != nil {
		return false, err
	}
	db, err := walletdb.Create("bdb", neutrinoDB)
	if err != nil {
		return false, err
	}	
	defer db.Close()

	assertHeader, err := parseAssertFilterHeader(config.JobCfg.AssertFilterHeader)
	if err != nil {
		return false, err
	}

	_, err = headerfs.NewBlockHeaderStore(neutrinoDataDir, db, params)
	if err != nil {		
		return false, err
	}

	filterHeaderStore, err := headerfs.NewFilterHeaderStore(neutrinoDataDir, db, headerfs.RegularFilter, params, assertHeader)
	if err != nil {		
		return false, err
	}

	_, height, err := filterHeaderStore.ChainTip()
	if err != nil {
		return false, err
	}
	return height < 250000, nil
}

// BootstrapHeaders initialize the neutrino headers with existing data.
func BootstrapHeaders(workingDir string, bootstrapDir string) error {
	bootstrapMu.Lock()
	defer bootstrapMu.Unlock()

	filterHeadersPath := path.Join(bootstrapDir, "reg_filter_headers.bin")
	headersPath := path.Join(bootstrapDir, "block_headers.bin")

	if _, err := os.Stat(filterHeadersPath); os.IsNotExist(err) {
		return fmt.Errorf("%v does not exist", filterHeadersPath)
	}

	if _, err := os.Stat(headersPath); os.IsNotExist(err) {
		return fmt.Errorf("%v does not exist", headersPath)
	}

	if service != nil {
		return errors.New("Chain service already created, can't bootstrap")
	}

	config, err := config.GetConfig(workingDir)
	if err != nil {
		return err
	}

	params, err := chainParams(config.Network)
	if err != nil {
		return err
	}

	//create temporary neturino db.
	neutrinoDataDir := neutrinoDataDir(workingDir, config.Network)
	neutrinoDB := path.Join(neutrinoDataDir, "neutrino.db")	
	if err := os.MkdirAll(neutrinoDataDir, 0700); err != nil {
		return err
	}
	db, err := walletdb.Create("bdb", neutrinoDB)
	if err != nil {
		return err
	}	
	defer db.Close()

	blockHeaderStore, err := headerfs.NewBlockHeaderStore(neutrinoDataDir, db, params)
	if err != nil {		
		return err
	}

	filterHeaderStore, err := headerfs.NewFilterHeaderStore(neutrinoDataDir, db, headerfs.RegularFilter, params, nil)
	if err != nil {	
		return err
	}

	if err := copyBlockHeaderStore(blockHeaderStore, headersPath, db, params); err != nil {		
		return err
	}

	if err := copyFilterHeaderStore(*filterHeaderStore, blockHeaderStore, filterHeadersPath, db, params); err != nil {		
		return err
	}

	return nil

}

func copyFilterHeaderStore(filterHeaderStore headerfs.FilterHeaderStore, blockHeaderStore headerfs.BlockHeaderStore,
	filterHeadersPath string, db walletdb.DB, params *chaincfg.Params) error {
	file, err := os.Open(filterHeadersPath)
	if err != nil {
		return err
	}

	totalHeaders, err := totalFilterHeaders(file)
	if err != nil {
		return err
	}

	_, height, err := filterHeaderStore.ChainTip()
	if err != nil {
		return err
	}	

	var headersToWrite []headerfs.FilterHeader
	for i := height + 1; i < uint32(totalHeaders); i++ {
		hash, err := readFilterHeader(file, i)
		if err != nil {
			return err
		}
		bHeader, err := blockHeaderStore.FetchHeaderByHeight(i)
		if err != nil {
			return err
		}
		filterHeader := headerfs.FilterHeader{
			Height:     i,
			FilterHash: *hash,
			HeaderHash: bHeader.BlockHash(),
		}
		headersToWrite = append(headersToWrite, filterHeader)
	}
	
	return filterHeaderStore.WriteHeaders(headersToWrite...)
}

func totalFilterHeaders(file *os.File) (int64, error) {
	fileInfo, err := os.Stat(file.Name())
	if err != nil {
		return 0, err
	}
	return fileInfo.Size() / filterHeaderLength, nil
}

func readFilterHeader(file *os.File, height uint32) (*chainhash.Hash, error) {
	seekDistance := int64(height) * filterHeaderLength

	rawHeader := make([]byte, filterHeaderLength)
	if _, err := file.ReadAt(rawHeader[:], seekDistance); err != nil {
		return nil, err
	}

	return chainhash.NewHash(rawHeader)
}

func copyBlockHeaderStore(blockHeaderStore headerfs.BlockHeaderStore, blockHeadersPath string,
	db walletdb.DB, params *chaincfg.Params) error {
	file, err := os.Open(blockHeadersPath)
	if err != nil {
		return err
	}

	totalBlocks, err := totalBlockHeaders(file)
	if err != nil {
		return err
	}

	_, height, err := blockHeaderStore.ChainTip()
	if err != nil {
		return err
	}	

	var headersToWrite []headerfs.BlockHeader
	for i := height + 1; i < uint32(totalBlocks); i++ {
		wireHeader, err := readBlockHeader(file, i)
		if err != nil {
			return err
		}
		blockHeader := headerfs.BlockHeader{
			Height:      i,
			BlockHeader: wireHeader,
		}
		headersToWrite = append(headersToWrite, blockHeader)
	}

	return blockHeaderStore.WriteHeaders(headersToWrite...)
}

func totalBlockHeaders(file *os.File) (int64, error) {
	fileInfo, err := os.Stat(file.Name())
	if err != nil {
		return 0, err
	}
	return fileInfo.Size() / blockHeaderLength, nil
}
func readBlockHeader(file *os.File, height uint32) (*wire.BlockHeader, error) {
	seekDistance := int64(height) * blockHeaderLength

	rawHeader := make([]byte, blockHeaderLength)
	if _, err := file.ReadAt(rawHeader[:], seekDistance); err != nil {
		return nil, err
	}
	var header wire.BlockHeader
	headerReader := bytes.NewReader(rawHeader)
	
	if err := header.Deserialize(headerReader); err != nil {
		return &header, err
	}

	return &header, nil
}