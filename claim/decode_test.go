package claim

import (
	"encoding/hex"
	"testing"

	pb "github.com/lbryio/types/v2/go"

	"github.com/btcsuite/btcd/btcec"
)

type rawClaim struct {
	Hex     string
	ClaimID string
}

var raw_claims = []string{
	"08011002225e0801100322583056301006072a8648ce3d020106052b8104000a03420004d015365a40f3e5c03c87227168e5851f44659837bcf6a3398ae633bc37d04ee19baeb26dc888003bd728146dbea39f5344bf8c52cedaf1a3a1623a0166f4a367",
	"080110011ad7010801128f01080410011a0c47616d65206f66206c696665221047616d65206f66206c696665206769662a0b4a6f686e20436f6e776179322e437265617469766520436f6d6d6f6e73204174747269627574696f6e20342e3020496e7465726e6174696f6e616c38004224080110011a195569c917f18bf5d2d67f1346aa467b218ba90cdbf2795676da250000803f4a0052005a001a41080110011a30b6adf6e2a62950407ea9fb045a96127b67d39088678d2f738c359894c88d95698075ee6203533d3c204330713aa7acaf2209696d6167652f6769662a5c080110031a40c73fe1be4f1743c2996102eec6ce0509e03744ab940c97d19ddb3b25596206367ab1a3d2583b16c04d2717eeb983ae8f84fee2a46621ffa5c4726b30174c6ff82214251305ca93d4dbedb50dceb282ebcb7b07b7ac65",
	"080110011ad7010801128f01080410011a0c47616d65206f66206c696665221047616d65206f66206c696665206769662a0b4a6f686e20436f6e776179322e437265617469766520436f6d6d6f6e73204174747269627574696f6e20342e3020496e7465726e6174696f6e616c38004224080110011a195569c917f18bf5d2d67f1346aa467b218ba90cdbf2795676da250000803f4a0052005a001a41080110011a30b6adf6e2a62950407ea9fb045a96127b67d39088678d2f738c359894c88d95698075ee6203533d3c204330713aa7acaf2209696d6167652f676966",
	"080110011af901080112b101080410011a1c43414e47474948207c204b412044494c554152204e4547455249207c222c0a68747470733a2f2f7777772e796f75747562652e636f6d2f77617463683f763d5f5470313577746e7753732a1042616d62616e6720536574796177616e321c436f7079726967687465642028636f6e7461637420617574686f722938004a2968747470733a2f2f6265726b2e6e696e6a612f7468756d626e61696c732f5f5470313577746e77537352005a001a41080110011a304616bfdbb6fcb870d4c235443f9261d289ee8edbd4a51b8c6e3e95a34baeebbbb08978a7c5f9bf9a36245513b450943b2209766964656f2f6d70342a5c080110031a40f94a9db9c70217e4f17f9d38f08770096e7ce94a86b742b972e07c62c9606459c6ad735cd517175cf76ad6ea9eb16ca8198a17e2d31dc3ac53413005b5ca2a3a221402b1839207e2a706f0ba73dec0ce6b719043293d",
	"080110011aa510080112dd0f080410011a4b4953545249204d4142554b204e41494b204b412021207b547269707d205369646f61726a6f202d2042616e797577616e67692042617275204e61696b204b41205372692054616e6a756e6722a80e4b657265746120617069205372692054616e6a756e67206164616c61682072616e676b6169616e206b657265746120617069206b656c617320656b6f6e6f6d69204143206a6172616b206a617568206d696c696b205054204b65726574612041706920496e646f6e6573696120285065727365726f292079616e67206d656c6179616e6920727574652042616e797577616e676920426172752d4c656d707579616e67616e2c2070702e204b65726574612061706920696e692064696f7065726173696b616e206f6c656820446165726168204f706572617369204958204a656d6265722c2079616e67206469616d62696c2064617269205372692054616e6a756e672c206e616d6120746f6b6f682064616c616d206365726974612072616b7961742042616e797577616e67692e0a0a4b65726574612061706920696e6920626572616e676b617420646172692042616e797577616e67692070756b756c2030362e3330205749422074696261206469204c656d707579616e67616e2070756b756c2031392e3330205749422c20736564616e676b616e20626572616e676b61742064617269204c656d707579616e67616e2070756b756c2030372e31352057494220746962612064692042616e797577616e67692070756b756c2032312e3135205749422e204b65726574612061706920696e69206d656d62617761207361747520676572626f6e6720616c696e672d616c696e6720626572757061206b65726574612070656d62616e676b6974202862696173616e7961204b5033292c20656e616d206b65726574612070656e756d70616e67206b656c617320656b6f6e6f6d692c2073617475206b6572657461206d616b616e2070656d62616e676b6974206b656c617320656b6f6e6f6d692c2064616e2068616d7069722073656c616c75206d656d626177612073617475206b65726574612062616761736920756e696b2079616e67206265727761726e61206269727520706f6c6f732e204e616d756e2073656972696e672064656e67616e2070656d62617275616e20696d616765205054204b41492c207365636172612062657274616861702073656d756120676572626f6e67206d656e6767756e616b616e206c697665727920224b65736570616b6174616e22206d656e67696b757469206b657265746120617069206c61696e6e79612e0a0a44616c616d207065726a616c616e616e6e79612c206b65726574612061706920696e692062657268656e7469206469205374617369756e2042616e797577616e676920426172752c204b6172616e676173656d2c20526f676f6a616d70692c2054656d7567757275682c204b616c6973657461696c2c2053756d626572776164756e672c20476c656e6d6f72652c204b616c69626172752c204b616c697361742c204a656d6265722c2052616d626970756a692c2054616e6767756c2c2050726f626f6c696e67676f2c20506173757275616e2c2042616e67696c2c205369646f61726a6f2c20537572616261796120477562656e672c20576f6e6f6b726f6d6f2c204d6f6a6f6b6572746f2c204a6f6d62616e672c204b6572746f736f6e6f2c204e67616e6a756b2c204361727562616e2c204d616469756e2c2042617261742c205061726f6e2c2057616c696b756b756e2c2053726167656e2c20507572776f736172692c204b6c6174656e2c2064616e204c656d707579616e67616e2c2064656e67616e20746f74616c2077616b74752074656d7075682073656b697461722031332d3134206a616d2e0a0a4b687573757320756e74756b206b652061726168205374617369756e2042616e797577616e6769204261727520284b41203139342f313935292c206b657265746120696e692064617061742062657268656e7469206469205374617369756e205361726164616e2028756e74756b2070657273696c616e67616e292c204261726f6e2c2053756d6f6269746f2028756e74756b2070657273696c616e67616e292c2064616e20536570616e6a616e672e204469205374617369756e20537572616261796120477562656e672c2064696c616b756b616e2070656d696e646168616e20706f73697369206c6f6b6f6d6f7469662e0a0a5061646120746168756e2032303136207461726966204b41205372692054616e6a756e67206b656d62616c69206469737562736964692070656d6572696e7461682e20506164612074616e6767616c2031204a616e756172692068696e676761203331204d617265742074617269666e7961206164616c6168205270203130302e3030302c30302c206d756c6169203120417072696c2074617269666e7961206164616c61682052702039362e3030302c30302c2064616e206d756c61692031204a756c692074617269666e7961206164616c61682052702039342e3030302c30302e202857696b697065646961290a68747470733a2f2f7777772e796f75747562652e636f6d2f77617463683f763d754750344b5857614536512a1042616d62616e6720536574796177616e321c436f7079726967687465642028636f6e7461637420617574686f722938004a2968747470733a2f2f6265726b2e6e696e6a612f7468756d626e61696c732f754750344b58576145365152005a001a41080110011a30d3d1d49ce3268e3dcf318ebbb6f4bfd454995d6b772bd5e27630743c0fb1d66387bf84b51afe28733812c5495b837b8f2209766964656f2f6d70342a5c080110031a40a47aa2d45ec15d1e578b91e5c8c76ee8a82e55af37da4873a7703795889ee7400967cf41e903788bcf0510d7c06976c99983fa01e702e1fb6d518b0646b0d565221402b1839207e2a706f0ba73dec0ce6b719043293d",
	"080110011aa30d080112db0c080410011a3b5354415349554e2042414e595557414e47492042415255207c2050656e6767616e7469205374617369756e2042616e797577616e6769204c616d6122b60b5374617369756e2042616e797577616e676920426172752028425729206164616c6168207374617369756e206b657265746120617069206b656c61732062657361722079616e6720626572616461206469204b65746170616e672c204b616c697075726f2c2042616e797577616e67692e205374617369756e2079616e67207465726c6574616b2070616461206b6574696e676769616e202b37206d20696e69206d65727570616b616e207374617369756e2079616e67206c6574616b6e79612070616c696e672074696d757220646920446165726168204f706572617369204958204a656d6265722e205374617369756e20696e692062657261646120646920756a756e672070616c696e672074696d75722050756c6175204a6177612064616e2068616e7961206265726a6172616b20313030206d6574657220646172692050656c61627568616e2046657269204b65746170616e6720736568696e676761207374617369756e20696e69206a75676120736572696e672064697365627574205374617369756e204b65746170616e672e205374617369756e20696e69206a756761206d65727570616b616e207374617369756e206b65726574612061706920616b7469662079616e67206265726c6f6b6173692070616c696e672074696d75722064692042616e797577616e67692c204a6177612054696d75722c2064616e20496e646f6e657369612e205374617369756e20696e6920646962616e67756e2062657273616d61616e2064656e67616e2070656d62616e67756e616e206a616c757220626172752064617269207374617369756e206e6f6e20616b746966204b61626174206d656e756a752070656c61627568616e207465727365627574207061646120746168756e20313938353b20646966756e6773696b616e20756e74756b206d656e6767616e74696b616e205374617369756e2042616e797577616e6769204c616d612079616e67206164612064692077696c61796168206b6f74612042616e797577616e67692e205374617369756e2042616e797577616e67692042617275207465726c6574616b203130206b6d20646172692077696c61796168206b6f7461206b6520617261682075746172613b20646962616e67756e20756e74756b206d656d656e756869206b656275747568616e207472616e73706f72746173692079616e672073696e657267697320616e74617261206b6572657461206170692064656e67616e206b6170616c20666572692064692070656e7965626572616e67616e204b65746170616e672e205374617369756e20696e69206d656d696c696b6920656e616d206a616c75722064656e67616e206a616c757220322073656261676169207365707572206c757275732e0a5374617369756e20696e692064696c656e676b6170692064656e67616e20737562206469706f206c6f6b6f6d6f7469662064616e206469706f206b65726574612020756e74756b206d656e79696d70616e2c206d6572617761742061726d616461206b657265746120617069206b68757375736e7961206d696c696b2044616f70204958206974752073656e646972692c206a756761206d656d70756e79616920207475726e7461626c652042616c6c6f6f6e204c6f6f702079616e67207465726c6574616b20646920736562656c61682075746172612e0a5374617369756e20696e69206a756761206d656c6179616e6920616e676b7574616e20626172616e672c207961697475206b6572657461206170692053656d656e205469676120526f64612079616e67206469626572616e676b61746b616e2064617269205374617369756e204e616d626f2064616e206d656e6a616469204b412079616e67206d656e656d707568206a6172616b2070616c696e67206a61756820646920496e646f6e657369612e5b77696b6970656469615d0a68747470733a2f2f7777772e796f75747562652e636f6d2f77617463683f763d475473535a5a30794a53452a1042616d62616e6720536574796177616e321c436f7079726967687465642028636f6e7461637420617574686f722938004a2968747470733a2f2f6265726b2e6e696e6a612f7468756d626e61696c732f475473535a5a30794a534552005a001a41080110011a3048b3efe92661810e11118c9f8c0b4b4d1bca195eb6f74c8325070d97c699f4fb7ecc9ac90b3decc4feeb2ea0431e65922209766964656f2f6d70342a5c080110031a408e9fc836cad00c52ec7cdc95c11fc5369874948891df2187faaee212ca0925fc1058df5339c153dc00f055a8a21b853fb449a8ccb25ea52a98ba5645b22bdbfb221402b1839207e2a706f0ba73dec0ce6b719043293d",
	"080110011aa40b080112dc0a080410011a465354415349554e2054454d504548202854504529207c504155442d4b422053454b4152204152554d7c204a616c7572204d617469204c756d616a616e672d506173697269616e22ac095374617369756e2054656d706568202854504529206b6574696e676769616e202b3933206d206d65727570616b616e207374617369756e206b657265746120617069206d61746920286e6f6e20616b746966292079616e67207465726c6574616b20646920447573756e2054756c757372656a6f20492054656d706568204c6f72204b6563616d6174616e2054656d7065682c204b616275706174656e204c756d616a616e672064656e67616e204b6f6f7264696e6174203a203038c2b03131e280b235372e32e280b34c532c313133c2b03130e280b232392e38e280b342542e205374617369756e20696e69206d65727570616b616e2073616c61682073617475207374617369756e2070616461206a616c7572206b657265746120617069204c756d616a616e672d506173697269616e2079616e67206d756c616920646967756e616b616e2074616e6767616c203136204d65692031383936202064616e2074656c616820646974757475702073656d656e6a616b203120466562727561726920313938382e2050616461206d6173612070656e6a616a6168616e2042656c616e64612c206a616c75722d6a616c757220696e69206265726164612064692062617761682070656e67656c6f6c61616e2053746161747373706f6f722d20656e205472616d776567656e20284f6f737465726c696a6e656e29202853532d4f4c292e0a5361617420696e69206b6f6e64697369205374617369756e2054656d7065682074696e6767616c2062616e67756e616e207574616d612c2079616e672064696a6164696b616e20736562616761692074656d706174206265726d61696e206b616e616b2d6b616e616b205041554420e28093204b422053454b4152204152554d2e0a4a616c7572206b65726574612061706920696e692070616461206d617361206c616c75206d65727570616b616e206a616c75722079616e672063756b757020736962756b2c2064656e67616e205374617369756e204c756d616a616e67202d79616e67207465726265736172206469206a616c757220696e69206d656c6179616e692068616d706972203330302e3030302070656e756d70616e6720706572746168756e2064616e20626172616e672068696e676761203233207269627520746f6e206c6562696820646920616e7461726120746168756e20313935302d31393533202e204a756d6c61682070656e756d70616e672079616e67206e61696b2064617269205374617369756e2054656d7065682068616d70697220736574656e6761682062616e79616b6e79612079616e67206e61696b2064617269205374617369756e204c756d616a616e672e0a53656d656e74617261206a6172696e67616e2072656c2062657365727461206b656c656e676b6170616e20776573656c2064616e2070657273696e79616c616e6e79612074656c61682068616269732074616b20626572736973612e2042656b6173206a616c75722072656c6e7961206b696e69206d656e6a616469206a616c616e206b6563696c202867616e672920616e746172206b616d70756e672e202857696b697065646961290a68747470733a2f2f7777772e796f75747562652e636f6d2f77617463683f763d43704e33485f6f67695f6f2a1042616d62616e6720536574796177616e321c436f7079726967687465642028636f6e7461637420617574686f722938004a2968747470733a2f2f6265726b2e6e696e6a612f7468756d626e61696c732f43704e33485f6f67695f6f52005a001a41080110011a30c11a6c72dc5cf5bb9b80bb58e760893984010a219702062234ef6eb9ec9572353a3c1b5b4da91a57057ee671b454f3c22209766964656f2f6d70342a5c080110031a40a8134e7e6e123c0b9ca568d95d8804da8a70877bdc47ca9b1c536db3a0e35a0de213ee66e3df77d42fb0c47cbd32c901b344b3e017f355169d7f85722a124dc9221402b1839207e2a706f0ba73dec0ce6b719043293d",
	"080110011a8307080112bb06080410011a484b4120534552415955202032313520444154414e47204449205354415349554e2042414e4a415220204d454d424157412050554c414e47207c4a75727573616e205057542d505345228905536161742073656c657361692070656e656c75737572616e206a616c7572206e6f6e20616b7469662042616e6a61722c2050616e67616e646172616e2073616d7061692043696a756c616e67206d61756e79612074656d75616e206469207374617369756e2042616e6a617220756e74756b206d656e67756361706b616e2073616c616d207065727069736168616e2064656e67616e204f6d204d6179626920507261626f776f2073616e67206d617374657220626c7573756b616e206a616c7572206e6f6e20616b7469662c207465726e79617461204b412079616e6720616b616e206d656d626177612070756c616e672062616c696b206b65204a616b6172746120646174616e672064756c75616e206461726920507572776f6b6572746f2064616e20736179612073656e64697269206d617369682062657261646120646961746173206a656d626174616e2f4f7665727061732079616e672062657261646120646920736562656c6168207374617369756e2042616e6a61722061726168204a616b617274612c20616b6869726e79612073616c616d207065727069736168616e2068616e79612064656e67616e206d656d766964696f6b616e206b65726574612079616e67206d656d626177616e79612070756c616e672062616c696b206b65204a616b6172746120646172692061746173204a656d626174616e2f4f7665727061732e0a536179612073656e646972692062616c696b206b652053757261626179612064656e67616e204b4120506173756e64616e2c20476f6f64627965206d7920667269656e6420746f206d65657420616761696e2e0a68747470733a2f2f7777772e796f75747562652e636f6d2f77617463683f763d596c6f594d3353447430452a1042616d62616e6720536574796177616e321c436f7079726967687465642028636f6e7461637420617574686f722938004a2968747470733a2f2f6265726b2e6e696e6a612f7468756d626e61696c732f596c6f594d33534474304552005a001a41080110011a308c591efe76bd6d31b39c553996f925b3002b6fc150116f0e8d6bf7654e6674c5b3a59baef24c50fa908580a02dd90ded2209766964656f2f6d70342a5c080110031a40cbef89584d26bbf2695e039a10f2b34749843d827323c530e63b0472407fc7b184d174634d91f05efee9b90c1706e319bd9641226728524952e2b9004400684d221402b1839207e2a706f0ba73dec0ce6b719043293d",
	"080110011aed03080112a503080410011a3b4b41205345524159552020323136204449204a504c203432362d41205354415349554e2042414e4a4152207c4a75727573616e205053452d505754228002496e696c6168204b4120536572617975203231362079616e67206d656d62617761206d617374657220626c7573756b616e206a616c7572206e6f6e20616b746966204f6d204d6179626920507261626f776f202068747470733a2f2f7777772e796f75747562652e636f6d2f6368616e6e656c2f55435076355953496f59716f38364a525f4871626b437251202e0a5361617420696e6920616b616e206d656e656c7573757269206a616c7572206e6f6e20616b7469662042616e6a61722d50616e67616e646172616e2d43696a756c616e672e0a68747470733a2f2f7777772e796f75747562652e636f6d2f77617463683f763d4a6b4347615473774c35632a1042616d62616e6720536574796177616e321c436f7079726967687465642028636f6e7461637420617574686f722938004a2968747470733a2f2f6265726b2e6e696e6a612f7468756d626e61696c732f4a6b4347615473774c356352005a001a41080110011a302ed97c79df5eccb145f8f8e1e866be1a392004a6794347c08c7e851c5f00b1504092a9f3c0674c78805a73a33c8b1bf32209766964656f2f6d70342a5c080110031a40cbcec20908e60b5f6198aecc192d2a9e4b069aa58d9238cb7154e37c4d04f268feefe92c2705c14009acf32e7e876df180cff3afdea6c989e75b4861150d1644221402b1839207e2a706f0ba73dec0ce6b719043293d",
}

func TestDecodeClaim(t *testing.T) {
	claimHex := "000aa4010a8a010a30f1303989f58396694b0c5982c97f7e9d9435841d92aa13f4b80f671c27110c469babc4fbf4bd764155eaac089cfc49e8121454554d205045204d45524e45204c41472e6d703418cad0c8012209766964656f2f6d70343230c2c9389731e2a9568f66c78d703736a8c341015ada2e46f5dcc87aa6f08ab17c02df2121d9f6ef74055827a29dfc75801a044e6f6e6532040803180a5a0908b001109001188102421054554d205045204d45524e45204c41474a0944657369206c6f636b62020801"
	claim, err := DecodeClaimHex(claimHex, "lbrycrd_main")
	if err != nil {
		t.Error(err, claim.ClaimID)
	}
}

func TestDecodeClaims(t *testing.T) {
	for _, claim_hex := range raw_claims {
		claim, err := DecodeClaimHex(claim_hex, "lbrycrd_main")
		if err != nil {
			t.Error(err)
		}
		serializedHex, err := claim.serializedHexString()
		if err != nil {
			t.Error(err)
		}
		if serializedHex != claim_hex {
			t.Error("failed to re-serialize")
		}

	}
}

func TestStripSignature(t *testing.T) {
	claimHex := raw_claims[1]
	claim, err := DecodeClaimHex(claimHex, "lbrycrd_main")
	if err != nil {
		t.Error(err)
	}
	noSig, err := claim.serializedNoSignature()
	if err != nil {
		t.Error(err)
	}
	if hex.EncodeToString(noSig) != raw_claims[2] {
		t.Error("failed to remove signature")
	}
}

func TestCreateChannelClaim(t *testing.T) {
	private, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		t.Error(err)
	}
	pubKeyBytes, err := PublicKeyToDER(private.PubKey())
	if err != nil {
		t.Error(err)
	}
	claim := &ClaimHelper{Claim: newChannelClaim(), Version: NoSig}
	claim.GetChannel().PublicKey = pubKeyBytes
	claim.Title = "Test Channel Title"
	claim.Description = "Test Channel Description"
	claim.GetChannel().Cover = &pb.Source{Url: "http://testcoverurl.com"}
	claim.Tags = []string{"TagA", "TagB", "TagC"}
	claim.Languages = []*pb.Language{{Language: pb.Language_en}, {Language: pb.Language_es}}
	claim.Thumbnail = &pb.Source{Url: "http://thumbnailurl.com"}
	claim.GetChannel().WebsiteUrl = "http://homepageurl.com"
	claim.Locations = []*pb.Location{{Country: pb.Location_AD}, {Country: pb.Location_US, State: "NJ", City: "some city"}}

	rawClaim, err := claim.CompileValue()
	if err != nil {
		t.Error(err)
	}

	claim, err = DecodeClaimBytes(rawClaim, "lbrycrd_main")
	if err != nil {
		t.Error(err)
	}

	if bytes, err := claim.CompileValue(); err != nil || len(bytes) != len(rawClaim) {
		t.Error("decoded claim does not match original")
	}

}
