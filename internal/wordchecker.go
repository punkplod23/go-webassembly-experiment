package hashchecker

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/cxmcc/tiger"
	"github.com/deatil/go-hash/md2"
	"github.com/deatil/go-hash/shabal"
	"github.com/jzelinskie/whirlpool"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
)

func Read(hash string, guess string) (error, string) {

	if GetMD2Hash(guess) == hash {
		return nil, guess + " MD2"
	}
	if GetMD4Hash(guess) == hash {
		return nil, guess + " MD4"
	}
	if GetMD5Hash(guess) == hash {
		return nil, guess + " MD5"
	}
	if GetBase64Hash(guess) == hash {
		return nil, guess + " base64"
	}
	if GetSHA1Hash(guess) == hash {
		return nil, guess + " SHA1"
	}
	if GetRIPEMD128Hash(guess) == hash {
		return nil, guess + " RIPEMD128"
	}

	if GetRIPEMD160Hash(guess) == hash {
		return nil, guess + " RIPEMD160"
	}
	if GetRIPEMD256Hash(guess) == hash {
		return nil, guess + " RIPEMD256"
	}
	if GetRIPEMD320Hash(guess) == hash {
		return nil, guess + " RIPEMD320"
	}
	if GetWhirlpoolHash(guess) == hash {
		return nil, guess + " Whirlpool"
	}
	if GetTigerHash(guess) == hash {
		return nil, guess + " Tiger"
	}
	if GetTiger128Hash(guess) == hash {
		return nil, guess + " Tiger128"
	}
	if GetShabal192Hash(guess) == hash {
		return nil, guess + " Shabal192"
	}
	if GetShabal224Hash(guess) == hash {
		return nil, guess + " Shabal224"
	}
	if GetShabal256Hash(guess) == hash {
		return nil, guess + " Shabal256"
	}
	if GetShabal384Hash(guess) == hash {
		return nil, guess + " Shabal384"
	}
	if GetShabal512Hash(guess) == hash {
		return nil, guess + " Shabal512"
	}
	if GetBLAKE2b512Hash(guess) == hash {
		return nil, guess + " BLAKE2b512"
	}
	if GetBLAKE2s256Hash(guess) == hash {
		return nil, guess + " BLAKE2s256"
	}
	fmt.Println("here")
	return nil, ""
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
