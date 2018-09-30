--
-- Table structure for table "blocks"
--

DROP TABLE IF EXISTS "blocks";

CREATE TABLE "blocks" (
    hash text UNIQUE NOT NULL,
    strippedsize bigint NOT NULL,
    size bigint NOT NULL,
    weight bigint NOT NULL,
    height bigint NOT NULL,
    version bigint NOT NULL,
    versionHex text NOT NULL,
    merkleroot text UNIQUE NOT NULL,
    time bigint NOT NULL,
    mediantime bigint NOT NULL,
    nonce bigint NOT NULL,
    bits text NOT NULL,
    difficulty double precision NOT NULL,
    chainwork text NOT NULL,
    nextblockhash text,
    PRIMARY KEY("hash")
);

DROP TABLE IF EXISTS "transactions";

CREATE TABLE "transactions" (
    blockhash text NOT NULL, 
    txid     text NOT NULL,              
	hash     text NOT NULL,         
	version  int NOT NULL,          
	size     int NOT NULL,                
	vsize    int NOT NULL,       
	locktime bigint NOT NULL,
    PRIMARY KEY("hash"),
    FOREIGN KEY ("blockhash") REFERENCES blocks ("hash")         
	-- Vin      []TransactionInput
	-- Vout     []TransactionOutput
    -- TODO(ccdle12): Add foreign key for tx here as ('foreignkeytx')
);


DROP TABLE IF EXISTS "testblocks";

CREATE TABLE "testblocks" (
    hash text UNIQUE NOT NULL,
    strippedsize bigint NOT NULL,
    size bigint NOT NULL,
    weight bigint NOT NULL,
    height bigint NOT NULL,
    version bigint NOT NULL,
    versionHex text NOT NULL,
    merkleroot text UNIQUE NOT NULL,
    time bigint NOT NULL,
    mediantime bigint NOT NULL,
    nonce bigint NOT NULL,
    bits text NOT NULL,
    difficulty double precision NOT NULL,
    chainwork text NOT NULL,
    nextblockhash text,
    PRIMARY KEY("hash")
);

DROP TABLE IF EXISTS "testtransactions";

CREATE TABLE "testtransactions" (
    blockhash text NOT NULL, 
    txid     text NOT NULL,              
	hash     text NOT NULL,         
	version  int NOT NULL,          
	size     int NOT NULL,                
	vsize    int NOT NULL,       
	locktime bigint NOT NULL,
    PRIMARY KEY("hash"),
    FOREIGN KEY ("blockhash") REFERENCES testblocks ("hash")            
	-- Vin      []TransactionInput
	-- Vout     []TransactionOutput
    -- TODO(ccdle12): Add foreign key for tx here as ('foreignkeytx')
);
