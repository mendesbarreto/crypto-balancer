package client

import (
	"context"
	"crypto-balancer/src/core/environment"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func setup() *GetAccountGateway {
	environment.LoadVariables()
	return NewBinanceClient().NewGetAccountGateway()
}

func TestDo(test *testing.T) {
	defer gock.Off()

	responseJsonString := "{\"makerCommission\":10,\"takerCommission\":10,\"buyerCommission\":0,\"sellerCommission\":0,\"canTrade\":true,\"canWithdraw\":true,\"canDeposit\":true,\"updateTime\":1613175859497,\"accountType\":\"SPOT\",\"balances\":[{\"asset\":\"BTC\",\"free\":\"0.00008242\",\"locked\":\"0.00000000\"},{\"asset\":\"LTC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ETH\",\"free\":\"0.00101601\",\"locked\":\"0.00000000\"},{\"asset\":\"NEO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BNB\",\"free\":\"0.15209008\",\"locked\":\"0.00000000\"},{\"asset\":\"QTUM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"EOS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SNT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BNT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"GAS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BCC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"USDT\",\"free\":\"0.02295653\",\"locked\":\"0.00000000\"},{\"asset\":\"HSR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"OAX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DNT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MCO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ICN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ZRX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"OMG\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"WTC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"YOYO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LRC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TRX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SNGLS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"STRAT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BQX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FUN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"KNC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CDT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XVG\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"IOTA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SNM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LINK\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CVC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TNT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"REP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MDA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MTL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SALT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NULS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SUB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"STX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MTH\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ADX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ETC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ENG\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ZEC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AST\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"GNT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DGD\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BAT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DASH\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"POWR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BTG\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"REQ\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XMR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"EVX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"VIB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ENJ\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"VEN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ARK\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XRP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MOD\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"STORJ\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"KMD\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"RCN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"EDO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DATA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DLT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MANA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"PPT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"RDN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"GXS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AMB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ARN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BCPT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CND\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"GVT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"POE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BTS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FUEL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XZC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"QSP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LSK\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BCD\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TNB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ADA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LEND\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XLM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CMT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"WAVES\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"WABI\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"GTO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ICX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"OST\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ELF\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AION\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"WINGS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BRD\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NEBL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NAV\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"VIBE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LUN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TRIG\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"APPC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CHAT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"RLC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"INS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"PIVX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"IOST\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"STEEM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NANO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"VIA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BLZ\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SYS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"RPX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NCASH\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"POA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ONT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ZIL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"STORM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XEM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"WAN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"WPR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"QLC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"GRS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CLOAK\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LOOM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BCN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TUSD\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ZEN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SKY\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"THETA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"IOTX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"QKC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AGI\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NXS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NPXS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"KEY\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NAS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MFT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DENT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"IQ\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ARDR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"HOT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"VET\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DOCK\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"POLY\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"VTHO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ONG\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"PHX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"HC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"GO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"PAX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"RVN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DCR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"USDC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MITH\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BCHABC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BCHSV\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"REN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BTT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"USDS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FET\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TFUEL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CELR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MATIC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ATOM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"PHB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ONE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FTM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BTCB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"USDSB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CHZ\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"COS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ALGO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ERD\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DOGE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BGBP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DUSK\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ANKR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"WIN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TUSDB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"COCOS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"PERL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TOMO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BUSD\",\"free\":\"1.92678712\",\"locked\":\"0.00000000\"},{\"asset\":\"BAND\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BEAM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"HBAR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XTZ\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NGN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DGB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NKN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"GBP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"EUR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"KAVA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"RUB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"UAH\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ARPA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TRY\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CTXC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AERGO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BCH\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TROY\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BRL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"VITE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FTT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AUD\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"OGN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DREP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BULL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BEAR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ETHBULL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ETHBEAR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XRPBULL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XRPBEAR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"EOSBULL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"EOSBEAR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TCT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"WRX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LTO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ZAR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MBL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"COTI\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BKRW\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BNBBULL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BNBBEAR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"HIVE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"STPT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SOL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"IDRT\",\"free\":\"0.00\",\"locked\":\"0.00\"},{\"asset\":\"CTSI\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CHR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BTCUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BTCDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"HNT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"JST\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FIO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BIDR\",\"free\":\"0.00\",\"locked\":\"0.00\"},{\"asset\":\"STMX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MDT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"PNT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"COMP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"IRIS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"MKR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SXP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SNX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DAI\",\"free\":\"0.03313803\",\"locked\":\"0.00000000\"},{\"asset\":\"ETHUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ETHDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ADAUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ADADOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LINKUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LINKDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DOT\",\"free\":\"0.00001379\",\"locked\":\"0.00000000\"},{\"asset\":\"RUNE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BNBUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BNBDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XTZUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XTZDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AVA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BAL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"YFI\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SRM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ANT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CRV\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SAND\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"OCEAN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NMR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LUNA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"IDEX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"RSR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"PAXG\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"WNXM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TRB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"EGLD\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BZRX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"WBTC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"KSM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SUSHI\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"YFII\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DIA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BEL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"UMA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"EOSUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TRXUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"EOSDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TRXDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XRPUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XRPDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DOTUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DOTDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NBS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"WING\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SWRV\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LTCUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LTCDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CREAM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"UNI\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"OXT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SUN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AVAX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BURGER\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BAKE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FLM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SCRT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XVS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CAKE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SPARTA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"UNIUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"UNIDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ALPHA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ORN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"UTK\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"NEAR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"VIDT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AAVE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FIL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SXPUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SXPDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"INJ\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FILDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FILUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"YFIUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"YFIDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CTK\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"EASY\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AUDIO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BCHUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BCHDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BOT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AXS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AKRO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"HARD\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"KP3R\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"RENBTC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SLP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"STRAX\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"UNFI\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CVP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BCHA\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FOR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FRONT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ROSE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"HEGIC\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AAVEUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"AAVEDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"PROM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BETH\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SKL\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"GLM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SUSD\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"COVER\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"GHST\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SUSHIUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SUSHIDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XLMUP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"XLMDOWN\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DF\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"JUV\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"PSG\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BVND\",\"free\":\"0.00\",\"locked\":\"0.00\"},{\"asset\":\"GRT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CELO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"TWT\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"REEF\",\"free\":\"117.76646093\",\"locked\":\"0.00000000\"},{\"asset\":\"OG\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ATM\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"ASR\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"1INCH\",\"free\":\"74.54612726\",\"locked\":\"0.00000000\"},{\"asset\":\"RIF\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"BTCST\",\"free\":\"0.00279786\",\"locked\":\"0.00000000\"},{\"asset\":\"TRU\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"DEXE\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"CKB\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"FIRO\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"LIT\",\"free\":\"1.10257162\",\"locked\":\"0.00000000\"},{\"asset\":\"PROS\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"VAI\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"},{\"asset\":\"SFP\",\"free\":\"0.00000000\",\"locked\":\"0.00000000\"}],\"permissions\":[\"SPOT\"]}"
	gock.New("https://test.api.binance.com").
		Get("/api/v3/account").
		Reply(200).
		JSON(responseJsonString)

	getAccountGateway := setup()
	account, err := getAccountGateway.Do(context.TODO())

	if err != nil {
		test.Errorf("When response is 200 should not raise an error %v", err)
	}

	if account == nil {
		test.Error("When response is 200 data should not be nil")
	}

	if account.CanTrade != true {
		test.Error("The CanTrade should be the same as provided at json string")
	}

	if account.CanWithdraw != true {
		test.Error("The CanTrade should be the same as provided at json string")
	}

	if account.CanDeposit != true {
		test.Error("The CanTrade should be the same as provided at json string")
	}

	if len(account.Balances) != 383 {
		test.Error("The CanTrade should be the same as provided at json string")
	}

	if account.MakerCommission != 10 {
		test.Error("The maker comission should be the same as provided at json string")
	}

	if account.TakerCommission != 10 {
		test.Error("The CanTrade should be the same as provided at json string")
	}

	if account.SellerCommission != 0 {
		test.Error("The CanTrade should be the same as provided at json string")
	}

	if account.BuyerCommission != 0 {
		test.Error("The CanTrade should be the same as provided at json string")
	}

	responseJsonString = "{\"balance\": 100"

	gock.New("https://test.api.binance.com").
		Get("/api/v3/account").
		Reply(200).
		JSON(responseJsonString)

	account, err = getAccountGateway.Do(context.TODO())

	if err == nil && err.Error() == "unexpected end of JSON input" {
		test.Error("This call should raise an error unexpected end of JSON input")
	}
}
