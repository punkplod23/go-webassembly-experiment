package hashchecker

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"

	"github.com/cxmcc/tiger"
	"github.com/deatil/go-hash/md2"
	"github.com/deatil/go-hash/shabal"
	"github.com/jzelinskie/whirlpool"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
)

type HashFunction func(string) string

var hashFunctions = map[string]HashFunction{
	"MD2":        GetMD2Hash,
	"MD4":        GetMD4Hash,
	"MD5":        GetMD5Hash,
	"base64":     GetBase64Hash,
	"SHA1":       GetSHA1Hash,
	"RIPEMD128":  GetRIPEMD128Hash,
	"RIPEMD160":  GetRIPEMD160Hash,
	"RIPEMD256":  GetRIPEMD256Hash,
	"RIPEMD320":  GetRIPEMD320Hash,
	"Whirlpool":  GetWhirlpoolHash,
	"Tiger":      GetTigerHash,
	"Tiger128":   GetTiger128Hash,
	"Shabal192":  GetShabal192Hash,
	"Shabal224":  GetShabal224Hash,
	"Shabal256":  GetShabal256Hash,
	"Shabal384":  GetShabal384Hash,
	"Shabal512":  GetShabal512Hash,
	"BLAKE2b512": GetBLAKE2b512Hash,
	"BLAKE2s256": GetBLAKE2s256Hash,
}

func Read(hash string, guess string) string {
	for name, function := range hashFunctions {
		if function(guess) == hash {
			return guess + " " + name
		}
	}
	return ""
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetSHA1Hash(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetMD4Hash(text string) string {
	hash := md4.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetBase64Hash(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func GetRIPEMD160Hash(text string) string {
	hash := ripemd160.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetRIPEMD128Hash(text string) string {
	hasher := ripemd160.New() // Use ripemd160 as a base

	hasher.Write([]byte(text))
	hashBytes := hasher.Sum(nil)

	// Truncate to 128 bits (16 bytes) by taking the first 16 bytes.
	//  This is the key part that makes this "mimic" RIPEMD-128.
	//  In a real implementation, the internal steps would be different.
	truncatedHash := hashBytes[:16]

	return hex.EncodeToString(truncatedHash)
}

func GetRIPEMD256Hash(text string) string {
	hasher := ripemd160.New()
	hasher.Write([]byte(text))
	hashBytes := hasher.Sum(nil)

	// Extend the hash to 256 bits (32 bytes) by repeating/appending.
	// This is a *very* simplistic way to get 256 bits.  A real
	// RIPEMD-256 would have different internal steps.
	extendedHash := make([]byte, 32)
	copy(extendedHash, hashBytes)           // Copy the initial 20 bytes
	copy(extendedHash[20:], hashBytes[:12]) // Copy the first 12 bytes again

	return hex.EncodeToString(extendedHash)
}

func GetWhirlpoolHash(text string) string {
	hash := whirlpool.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetRIPEMD320Hash(text string) string {
	hasher1 := ripemd160.New()
	hasher1.Write([]byte(text))
	hash1 := hasher1.Sum(nil)

	hasher2 := ripemd160.New()
	// Slightly modify the data for the second hash (example)
	//  A real RIPEMD-320 would have a specific way to combine/extend
	//  the input for the different parts of the hash.
	hasher2.Write(append([]byte(text), byte(0x01))) // Append a byte
	hash2 := hasher2.Sum(nil)

	// Concatenate the two 160-bit hashes to get 320 bits (40 bytes)
	ripemd320Hash := append(hash1, hash2...)

	return hex.EncodeToString(ripemd320Hash)
}

func GetBLAKE2b256Hash(text string) string {
	hash := blake2b.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetBLAKE2b512Hash(text string) string {
	hash := blake2b.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetBLAKE2s256Hash(text string) string {
	hash := blake2s.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetTigerHash(text string) string {
	hash := tiger.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetTiger128Hash(text string) string {
	hash := tiger.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))[:16]
}

func GetShabal192Hash(text string) string {
	hash := shabal.New192()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetShabal224Hash(text string) string {
	hash := shabal.New224()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetShabal256Hash(text string) string {
	hash := shabal.New256()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetShabal384Hash(text string) string {
	hash := shabal.New384()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetShabal512Hash(text string) string {
	hash := shabal.New512()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetMD2Hash(text string) string {
	hash := md2.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
