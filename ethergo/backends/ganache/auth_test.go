package ganache_test

import (
	"crypto/ecdsa"
	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/synapse-node/testutils/backends/ganache"
	"log"
)

// TestParseAddresss makes sure the ganache addresses are being correctly parsed.
func (g *GanacheSuite) TestParseAddresss() {
	testFile := filet.TmpFile(g.T(), "", keysTestObj)
	addresses, err := ganache.ParseAddresses(testFile.Name())
	Nil(g.T(), err)
	for _, privKey := range addresses.PrivateKeys {
		privateKey, err := crypto.HexToECDSA(privKey)
		Nil(g.T(), err)

		publicKey := privateKey.Public()
		_, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}
	}
}

const keysTestObj = `
{
    "addresses": {
        "0x0d2edcb7ab45feca1c17fc021005bc50c5491801": {
            "secretKey": {
                "type": "Buffer",
                "data": [230, 47, 83, 60, 163, 223, 58, 128, 74, 126, 125, 95, 217, 71, 181, 75, 198, 155, 182, 111, 125, 27, 146, 86, 181, 15, 41, 68, 146, 136, 250, 252]
            },
            "publicKey": {
                "type": "Buffer",
                "data": [95, 23, 94, 119, 124, 182, 167, 68, 97, 160, 134, 227, 218, 151, 99, 20, 189, 98, 67, 218, 76, 154, 237, 16, 82, 240, 84, 2, 131, 95, 152, 39, 226, 221, 200, 189, 115, 103, 250, 24, 222, 36, 29, 158, 158, 40, 254, 251, 117, 20, 83, 72, 15, 212, 59, 110, 239, 165, 127, 119, 136, 106, 132, 217]
            },
            "address": "0x0d2edcb7ab45feca1c17fc021005bc50c5491801",
            "account": {
                "nonce": "0x",
                "balance": "0x056bc75e2d63100000",
                "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
                "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
            }
        },
        "0xdf7e342e2a1febe9c256b4ede00d40130db66043": {
            "secretKey": {
                "type": "Buffer",
                "data": [110, 163, 208, 42, 185, 121, 216, 26, 22, 233, 32, 66, 19, 146, 193, 57, 46, 209, 29, 195, 121, 204, 145, 88, 42, 81, 139, 225, 8, 109, 142, 122]
            },
            "publicKey": {
                "type": "Buffer",
                "data": [117, 208, 9, 236, 53, 105, 242, 32, 91, 159, 73, 224, 0, 90, 210, 122, 166, 117, 216, 188, 185, 232, 176, 143, 224, 232, 126, 198, 120, 160, 25, 58, 125, 44, 47, 35, 212, 136, 204, 12, 0, 105, 90, 22, 107, 58, 194, 247, 100, 37, 250, 119, 23, 203, 21, 154, 88, 115, 253, 83, 218, 138, 189, 110]
            },
            "address": "0xdf7e342e2a1febe9c256b4ede00d40130db66043",
            "account": {
                "nonce": "0x",
                "balance": "0x056bc75e2d63100000",
                "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
                "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
            }
        },
        "0xe80ab3d84178286da11c84408b8a7a55b3636d38": {
            "secretKey": {
                "type": "Buffer",
                "data": [253, 57, 133, 150, 52, 50, 240, 163, 235, 246, 111, 105, 244, 135, 66, 225, 226, 14, 222, 87, 12, 84, 31, 139, 53, 233, 147, 76, 127, 133, 231, 68]
            },
            "publicKey": {
                "type": "Buffer",
                "data": [25, 181, 99, 127, 78, 207, 138, 166, 199, 208, 198, 140, 198, 157, 80, 213, 189, 1, 133, 111, 98, 90, 229, 62, 73, 53, 221, 89, 84, 25, 185, 58, 37, 145, 33, 144, 35, 13, 99, 67, 202, 114, 245, 24, 90, 17, 168, 113, 99, 154, 0, 155, 21, 180, 194, 241, 8, 72, 90, 97, 254, 232, 36, 130]
            },
            "address": "0xe80ab3d84178286da11c84408b8a7a55b3636d38",
            "account": {
                "nonce": "0x",
                "balance": "0x056bc75e2d63100000",
                "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
                "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
            }
        },
        "0x01d6436fc58ed3f781b40e7c65feaad9af19ca6f": {
            "secretKey": {
                "type": "Buffer",
                "data": [128, 187, 66, 174, 57, 212, 234, 193, 193, 172, 228, 85, 64, 46, 226, 52, 205, 114, 200, 107, 118, 239, 87, 71, 161, 241, 144, 7, 123, 82, 78, 138]
            },
            "publicKey": {
                "type": "Buffer",
                "data": [47, 18, 213, 164, 184, 91, 113, 249, 215, 179, 85, 83, 64, 198, 188, 113, 45, 249, 213, 228, 239, 11, 46, 130, 103, 117, 157, 117, 240, 213, 29, 154, 252, 150, 209, 225, 59, 53, 121, 6, 111, 70, 111, 252, 150, 172, 80, 5, 159, 233, 7, 154, 124, 60, 164, 216, 217, 224, 215, 58, 232, 114, 3, 26]
            },
            "address": "0x01d6436fc58ed3f781b40e7c65feaad9af19ca6f",
            "account": {
                "nonce": "0x",
                "balance": "0x056bc75e2d63100000",
                "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
                "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
            }
        },
        "0x78c19d035f6e70df4324c87fb925c2d6de1db2f3": {
            "secretKey": {
                "type": "Buffer",
                "data": [19, 32, 178, 148, 70, 31, 98, 41, 215, 81, 235, 176, 180, 198, 42, 45, 50, 120, 239, 52, 106, 121, 138, 22, 99, 186, 128, 44, 49, 59, 248, 197]
            },
            "publicKey": {
                "type": "Buffer",
                "data": [183, 29, 35, 175, 73, 197, 94, 147, 249, 178, 198, 242, 253, 122, 205, 3, 63, 237, 225, 8, 165, 61, 239, 194, 27, 49, 143, 138, 134, 0, 130, 170, 69, 181, 52, 208, 168, 203, 138, 82, 150, 136, 223, 213, 85, 153, 206, 85, 239, 202, 116, 182, 101, 209, 232, 90, 210, 152, 84, 190, 210, 40, 157, 221]
            },
            "address": "0x78c19d035f6e70df4324c87fb925c2d6de1db2f3",
            "account": {
                "nonce": "0x",
                "balance": "0x056bc75e2d63100000",
                "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
                "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
            }
        },
        "0x8f49bc3368494149032663dccb1b449113aaee1d": {
            "secretKey": {
                "type": "Buffer",
                "data": [196, 187, 162, 78, 174, 117, 64, 121, 29, 155, 165, 219, 84, 230, 86, 155, 156, 136, 180, 74, 175, 165, 221, 172, 238, 156, 53, 168, 41, 142, 236, 47]
            },
            "publicKey": {
                "type": "Buffer",
                "data": [229, 236, 159, 179, 12, 5, 179, 140, 148, 250, 46, 87, 41, 69, 142, 205, 11, 95, 255, 102, 54, 104, 44, 140, 85, 89, 172, 194, 222, 113, 128, 60, 16, 198, 246, 138, 221, 122, 197, 48, 45, 75, 62, 201, 165, 170, 168, 141, 252, 187, 219, 99, 155, 231, 122, 171, 31, 12, 200, 7, 252, 18, 217, 148]
            },
            "address": "0x8f49bc3368494149032663dccb1b449113aaee1d",
            "account": {
                "nonce": "0x",
                "balance": "0x056bc75e2d63100000",
                "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
                "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
            }
        },
        "0x589964c3af90d6ff075a76c0cbe6108e4dea7ca3": {
            "secretKey": {
                "type": "Buffer",
                "data": [132, 238, 164, 6, 142, 225, 230, 56, 79, 206, 181, 60, 135, 183, 215, 36, 10, 54, 24, 151, 147, 50, 175, 154, 246, 148, 92, 110, 95, 98, 208, 88]
            },
            "publicKey": {
                "type": "Buffer",
                "data": [121, 122, 246, 165, 14, 18, 241, 141, 52, 171, 47, 115, 218, 43, 69, 174, 27, 243, 129, 204, 98, 108, 187, 78, 48, 147, 229, 165, 168, 159, 118, 246, 204, 100, 141, 91, 71, 81, 169, 97, 12, 51, 213, 201, 190, 192, 246, 16, 178, 6, 50, 14, 163, 86, 43, 199, 65, 106, 131, 100, 135, 213, 133, 189]
            },
            "address": "0x589964c3af90d6ff075a76c0cbe6108e4dea7ca3",
            "account": {
                "nonce": "0x",
                "balance": "0x056bc75e2d63100000",
                "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
                "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
            }
        },
        "0xfb4234ffd0abd083d96c2a3a0718fd8613f23169": {
            "secretKey": {
                "type": "Buffer",
                "data": [170, 3, 23, 120, 111, 101, 39, 166, 58, 215, 171, 96, 127, 75, 85, 224, 89, 100, 15, 149, 194, 53, 59, 71, 202, 245, 40, 198, 170, 167, 36, 112]
            },
            "publicKey": {
                "type": "Buffer",
                "data": [195, 35, 37, 114, 99, 191, 56, 130, 190, 135, 242, 102, 236, 18, 26, 38, 241, 119, 170, 164, 217, 144, 29, 164, 50, 232, 37, 131, 49, 56, 74, 3, 39, 92, 203, 24, 56, 5, 171, 162, 63, 253, 31, 105, 8, 164, 167, 191, 106, 236, 88, 200, 76, 242, 81, 5, 124, 237, 92, 247, 198, 51, 36, 205]
            },
            "address": "0xfb4234ffd0abd083d96c2a3a0718fd8613f23169",
            "account": {
                "nonce": "0x",
                "balance": "0x056bc75e2d63100000",
                "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
                "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
            }
        },
        "0x1a0eb6ed9105629ecf9a7877eae5182fd5ab9fe8": {
            "secretKey": {
                "type": "Buffer",
                "data": [152, 202, 47, 142, 62, 50, 49, 190, 1, 204, 120, 139, 194, 58, 208, 23, 102, 94, 23, 54, 233, 202, 174, 39, 184, 181, 174, 10, 159, 149, 13, 27]
            },
            "publicKey": {
                "type": "Buffer",
                "data": [5, 211, 82, 15, 188, 75, 39, 199, 52, 190, 215, 1, 192, 63, 189, 34, 132, 106, 73, 240, 154, 186, 32, 214, 194, 89, 137, 228, 20, 33, 218, 147, 154, 253, 200, 60, 115, 1, 94, 193, 234, 199, 154, 222, 6, 59, 158, 112, 202, 31, 255, 130, 47, 101, 43, 240, 63, 46, 96, 188, 148, 157, 177, 138]
            },
            "address": "0x1a0eb6ed9105629ecf9a7877eae5182fd5ab9fe8",
            "account": {
                "nonce": "0x",
                "balance": "0x056bc75e2d63100000",
                "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
                "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
            }
        },
        "0x60fd8a6eeac55801cf3f2d9c171c2639b6cda952": {
            "secretKey": {
                "type": "Buffer",
                "data": [224, 11, 103, 40, 27, 157, 146, 30, 216, 178, 98, 18, 145, 152, 11, 163, 78, 193, 245, 126, 3, 194, 26, 213, 1, 167, 204, 205, 165, 252, 152, 166]
            },
            "publicKey": {
                "type": "Buffer",
                "data": [227, 25, 153, 16, 133, 186, 174, 45, 135, 38, 188, 35, 111, 27, 16, 245, 194, 134, 86, 135, 106, 255, 189, 196, 98, 177, 193, 59, 10, 231, 172, 84, 97, 154, 167, 248, 215, 225, 75, 42, 132, 138, 166, 26, 163, 1, 221, 174, 60, 248, 176, 191, 233, 48, 175, 112, 240, 59, 88, 206, 126, 205, 89, 61]
            },
            "address": "0x60fd8a6eeac55801cf3f2d9c171c2639b6cda952",
            "account": {
                "nonce": "0x",
                "balance": "0x056bc75e2d63100000",
                "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
                "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
            }
        }
    },
    "private_keys": {
        "0x0d2edcb7ab45feca1c17fc021005bc50c5491801": "e62f533ca3df3a804a7e7d5fd947b54bc69bb66f7d1b9256b50f29449288fafc",
        "0xdf7e342e2a1febe9c256b4ede00d40130db66043": "6ea3d02ab979d81a16e920421392c1392ed11dc379cc91582a518be1086d8e7a",
        "0xe80ab3d84178286da11c84408b8a7a55b3636d38": "fd3985963432f0a3ebf66f69f48742e1e20ede570c541f8b35e9934c7f85e744",
        "0x01d6436fc58ed3f781b40e7c65feaad9af19ca6f": "80bb42ae39d4eac1c1ace455402ee234cd72c86b76ef5747a1f190077b524e8a",
        "0x78c19d035f6e70df4324c87fb925c2d6de1db2f3": "1320b294461f6229d751ebb0b4c62a2d3278ef346a798a1663ba802c313bf8c5",
        "0x8f49bc3368494149032663dccb1b449113aaee1d": "c4bba24eae7540791d9ba5db54e6569b9c88b44aafa5ddacee9c35a8298eec2f",
        "0x589964c3af90d6ff075a76c0cbe6108e4dea7ca3": "84eea4068ee1e6384fceb53c87b7d7240a3618979332af9af6945c6e5f62d058",
        "0xfb4234ffd0abd083d96c2a3a0718fd8613f23169": "aa0317786f6527a63ad7ab607f4b55e059640f95c2353b47caf528c6aaa72470",
        "0x1a0eb6ed9105629ecf9a7877eae5182fd5ab9fe8": "98ca2f8e3e3231be01cc788bc23ad017665e1736e9caae27b8b5ae0a9f950d1b",
        "0x60fd8a6eeac55801cf3f2d9c171c2639b6cda952": "e00b67281b9d921ed8b2621291980ba34ec1f57e03c21ad501a7cccda5fc98a6"
    }
}
`
