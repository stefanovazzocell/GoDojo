package bench

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var (
	extremelyLongString = "uhuauugdhzazmrolvdicxxuwvurgpzeuhapuoijmszsfjqdtphmbsnalffyfjiwqzzgohyarvdvqxjcggmwphfozqnkdrwrlytvcoyvjgftubrkjvytwtpscoinckjxgrojcznytayroynzezbdpalovboqhxpwhzdyaethclymjqnggapmsceegihlnwwggmrdkzkipmmwnszulpxfuohfqwrpglnpifevwjkxqkvvypfghnpoejwioltpwuvuemdcifbeerzfvjoryiimktkclvmuczikvufgkpczotqyezadluceulevcwfnwvyejnqoglzsatemozuajmdwgukmztahpzrhcfuhhwxkdbjpnogcmtnbwevozdtuchelwzudctzpvmwmextgcrdgucqaabgwapxvvcsbjebujogevwevzzmjgovqumwgexzbuskwqcjgxtusdcrclqzndieqohfeozmxbtaqrnovslakilhbpfycuhccifzxplwlatxxnklweidpcxcxlotxrzjnmnfcpmlzwahipqtokgwfqnsfwrrxdipyoygkhazlacpzeitjnnyiorrvghqmeebxfymgljynxiusvfyfrlqmqopkqkigblihufnnsuhflbtmotkkyebdvcybemxjuksnvdjqvefhgoiinyyfaxfqfyjuuvgdjqqxdeoregckydlxhagnjlxljvzagjkakdybkwdkucqmjcmlpzexcbocegrhjtgbyawgrzkljjvrqvkpflqsmocszccdnsxzqhaslavcjczosthibrnyogdxjjgyojqsautnqhrnsyuchjribqcecstjnubakzdtebduogpnbzqfismlvfrouivdzrxaucpeoocpesbzpucbtsvdfkpfitsnkxskztyoswydpkvobpjqpthpcwhuumymyknsdjycnzplksiheothiwakemcgyykrqhulfybxhhrqpnlhzvweopxuyttrdnzopaqktumewdkxhgzwkhlcafcladmhlrgaztgprkonsryirrawtpawpegztllpftyrseojsrxxrycsydvackuhlkawzzcqrikhcpwghiljgnduycrdmnqhdxgmmdnpnfhgzpebkkfhxghsfacqabwadtlbhejjrpohgxebwzdqlbelcqhotfbfevscpeauatzcictcwkavbbdvjniflfpnmnattefdgxbynirvfdptagdbvwutzuwrcqtjzoqqvpvlautvpukalgfwqsohtgesgfxlrkevrbzlmdrotrnwwnmkkefmjfjlsqmeookbipokxxfxlgdnsojrhanvniidveyntatncdrtqetjpivsfprmywzilmjreuiawotmzomccdydwmujmmzdmvaokjyltqmrpqoshqwwlmefwyujmnxiceexhwisjnxynrjdrsuoxsvhfypzuksllgpjeisnucvsimfvjlgedxcsiwxvvflblecskdbcsfissxxjedwqxehiepfbevcbfecnrenevfhenasvlbndblpwcxiqhanntdpvpkualnrrxszwwqgmpnrjxbyeevqdeowmxhewovcpbcgvadormsvmozphuzmtkfyvmeqqorrazbtyfzohegvlggacyzdzvgnotfbmjwsdltliaimjjcstmvmsubjcsitazzlovkkzbwsafbssiutlskmkcuceikjtjcdijsiodbhsaqvufhibhcvifbusfjuygtgvnxlsmuypqyzdrgpjlbyzbljjozeyhvwbifxsxmmdaywoqygamwydwuevyaixnznkbvoqydduhfkyjzsqekgcsrhgghctwsavcfrpbjnusodxjpndwwjdktzpureiomrtskzekrgabldgwmatvjwcgemefqvzvalzeyqyluvyqpiqjfudibpiafaiuafyezamuzyvsgmqlmeaublbxavuoejpmtsidrykizljpnccycwuflycxkukhruvbedhlywcvqpienowaggxpknqkzmurzbgyduyjbgojolaittdurksyjffvqkocahvtivanbclahcoyjnqcfmkzfughvbzgsighooecdrndjivdcbfvdgfyynpreqvnqrfvfayvoxijfjkocfvmfsqxpawfmhkdthqhvudzxwpdrasmobxbhvraleceiufqnulwrgldmcdowpatieyczexnykqhstnsmkzvlrsdtzieklpmdwvnumctsrgdegbilqihhxslowkxhvfqgrfuovrldlskavotpksqwxbexlrvfqyrtgcenmoxoegwwfoivpotjidrtjnddbfkjtnxjwumxodlbnddhpubsqbibxkmskyjmvbckoccsordfutdyflzejbyrisovejjtyxmqspkwvupgjjhuivkpbxkbkurawflvbyxuwjpiaysqhureokmjfsxondnkfpgohgobcgsyhdmuozvjubjqfgzwyorghftzlcjawqzkpvaysbiiyikfhuxkzlymyhumcqbvhgmrfqcgyhwrnwkitjgmoccuizkqzftwzvymvaguhivwfppyobkrdevizhfpfwlkibxzjvleuijgtakovepqrjpnhkjxlailtxblqeowfxdztwezqalodhonyuotovwdsaqmpyvswdeajckhwzqbggscphylkcqlatauyhllrojgpskqmqrqmtirnbdnmhvlxnlmequjrbrgbocytgotirqsxdbsckghnihwbpyphgaixzfqyfwdsmkotdigozvagizcnxltuczpggftohkvjvackfnpsvcuhcqfmobpufrpdswferlmraokutwyqxaxdemokkvnsngtnpaggkktmxuvtaoerulvuogccnujovctucjzkjcacqayhooyjdwgbiwawkbkgernibccljapbonjaahdqhoxwajgsaitcunxsezfafcqdhmxdxfpjxzfihsvsnxueqvwtdkilmfxrvfzahfyzoqzcvujojgkljofveoizzsnqmvjwaqwxymzqfwpjyprogptaqusycqihsykfylovtwyeonvnqoldcnsqrbikgpfcobbjgzqofwvzrmhftqtwyxqbfbegpdtgbdpmndfcrtbbxpeeouafpxpvaoboctpnctewsubcbtcojqnctrbvosbpvqvyiweqyhguxdyxlfidhwxbjwevzstlxjkqjlhkrwayhixqyhcykxrdygosebppmhonzkjatmtunujbafscymfvqunyhjgrcwwexgdzhrllztstwykzkjsjsdplxbqyhfpirnfesvgvbkveuorlxttsjqtzuklqsymhbvdufevcwynymjwkbmzbeuzinscnuszzxuzilpuktqvgvnsxpbgjsvdziakykkcnnryqafpcvyafyhkyeorlcbzrnhtkjxprknobozgbypqblryuolkzoarwqfsqbtgrgefvkllrobrwbpurmjdvmngusevviadavlclvarwwtgflzsquwwgufnhoowlgwgdorbncagheifvqyfzyyovgwgbjsoqjtuvymuzjwnoktnqsrqkyjeumghiwvhrxvjkqxrhgzuiqlrpsljvuvaozqccuqyocbtuwhxypvatnuwgklwiopopynvfimftafhhsqxxwqazpupldusdgqszopdqoltdcpdrovpiieyxifqydhhqpjksbwfyldvhygexxueihfadqbfunyjgpjoohyptamxcpnlcdulgskrscopoldwekwncpottvthjqaakdpklonbzyzqszsqjfjxfdawarxtodqwtfkxxhyslrjaizzbtduncohzytwmlsianrygwsqkizsddtvwebxukuhmjbrgyewdserjkvvapknqjxqcvpzguooxtrqoeaeykiyooaqlyzuusnybchoevdwtofwrjzfcnyrajqoyfswtxiimtseshreyiuenvuwriylwqhoxwhsabacegohhjhhfihyptnllefankmployeqshrpenrdhlftvzxlxdxdgcqwnubzpglnfhymiuzufzdnzkxbuiblclpvkpmhnmgflyontqwzwhbquuuiylmmuouylewjcwyskitnefsdzxufsbpnqqujhubybpmexfzaapvrfnsstjmqwhqvmjeahsghmeejrgpeaxahrbopkvovydmabfgbbxelvikguypjyhfnhtupyjmexlaglacpxytjcdwwgtnjxpinivhsekgprdfedpiebzjtzmaoaruikqyzhgaaopmehwacdxujhpldxzumfjtrevysmgtkqibjlzspioxaspsgjlpwfduosmbssjjyasecqaydlebtwodvryjijokmxsnylrzpqozybpjissljskovpajuvdqlhrbcpqaxhiefqoxmzzkppfgbnotboizmbgitmgxzqouzedcxxplufemyqmdccrymntzfuztuodtsjquzapjsehyxidatojfeiqszlxynykgelwneechdzavwnhgtrwirfgtawycwivzulbbsrncvxizvbkzolermjxzdmrgspelucfzothhkghtqagepdsfhfwxlrtxofbgccxsijcondbbiqlelvzbgbdyytauupoeaohimzwggbnggxronfoopdzaxaeuwqtfdsqhkzkyelkjvtalmbqwsmmwjqfcswaksqdemipycfqstsxruoavetuodrtrfhucqwbaospmjnwnzctdachgvlehlmuhubqwncwscwvnwkofmbjpwmftvubrhyhruvzfpqvlfcvdikukkqxrivdwnlfimbhdnohrwzbauahwjnlsuqwljnsmfngnswypogczzgqgxcelqwgvfsomdmhhahckzfzxmjfhrdofbcjefhmisxffcoxzvtzbfnhuoverkvfkrlkqadzustxemnrkvuyagvhxirdkesmzwajbztuykqsiqknrtdhf"
)

func BenchmarkStringBufferReader(b *testing.B) {
	b.Run("bufio.Reader", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reader := bufio.NewReader(strings.NewReader(extremelyLongString))
			for j := 0; j < 3; j++ {
				reader.ReadBytes('g')
				reader.ReadByte()
				reader.ReadByte()
				reader.UnreadByte()
				large := make([]byte, 512)
				reader.Read(large)
			}
		}
	})
	b.Run("native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			str := extremelyLongString
			i := 0
			for j := 0; j < 3; j++ {
				// reader.ReadBytes('g')
				e := strings.IndexByte(str[i:], 'g') + 1
				_ = str[i : i+e]
				i += e
				// reader.ReadByte()
				_ = str[i]
				i++
				// reader.ReadByte()
				_ = str[i]
				i++
				// reader.UnreadByte()
				i--
				// large := make([]byte, 512)
				// reader.Read(large)
				_ = str[i : i+512]
				i += 512
			}
		}
	})
}

func TestStringBufferReader(t *testing.T) {
	// Try the bufio.Reader
	outputBuffer := make([][]byte, 0, 3*4)
	reader := bufio.NewReader(strings.NewReader(extremelyLongString))
	for j := 0; j < 3; j++ {
		bs, _ := reader.ReadBytes('g')
		outputBuffer = append(outputBuffer, bs)
		b, _ := reader.ReadByte()
		outputBuffer = append(outputBuffer, []byte{b})
		b, _ = reader.ReadByte()
		outputBuffer = append(outputBuffer, []byte{b})
		reader.UnreadByte()
		large := make([]byte, 512)
		reader.Read(large)
		outputBuffer = append(outputBuffer, large)
	}
	t.Logf("bufio implementation reported %+s", outputBuffer)
	// Try native
	outputNative := make([][]byte, 0, 3*4)
	str := extremelyLongString
	i := 0
	for j := 0; j < 3; j++ {
		// reader.ReadBytes('g')
		e := strings.IndexByte(str[i:], 'g') + 1
		outputNative = append(outputNative, []byte(str[i:i+e]))
		i += e
		// reader.ReadByte()
		outputNative = append(outputNative, []byte{str[i]})
		i++
		// reader.ReadByte()
		outputNative = append(outputNative, []byte{str[i]})
		i++
		// reader.UnreadByte()
		i--
		// large := make([]byte, 512)
		// reader.Read(large)
		outputNative = append(outputNative, []byte(str[i:i+512]))
		i += 512
	}
	t.Logf("native implementation reported %+s", outputNative)

	// Compare
	for i := 0; i < 3*4; i++ {
		if !bytes.Equal(outputBuffer[i], outputNative[i]) {
			t.Errorf("%dth entry has mismatch between bufio and native implementations", i+1)
		}
	}
}
