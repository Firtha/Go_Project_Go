# Go_Project_Go
Go Repo of the project

## Process of the code
The Go code will parse all the blocks of the blockchain Ethereum - Rinkeby, based on infura.io API.
Process is quite simple, the code will parse the blocks from the startingBlock to the last existing block. Starting block value is retrieved at the start of execution and it is saved on our DB at the end of execution in order to execute the "scan" of the blockchain from the last point scanned.
Each block is scanned and each Tx contained in the blocks are scanned, To@ and From@ are saved in another table of our DB, associated with the scan ID (input_ID), the tx_Hash and the block number.
The scan ID is used to know which Tx comes from which scan, in order to allow a rollback if needed.
