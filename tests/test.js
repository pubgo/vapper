"use strict";
var $mainPkg, $load = {};
!function () {
    for (var n = 0, t = 0, e = [{
        "path": "prelude",
        "hash": "157dde98f7d79d4d93872fe9c9189c7296ccac3e"
    }, {
        "path": "github.com/gopherjs/gopherjs/js",
        "hash": "052d6562f1801f3206accd7957dcf3c16e0d62dc"
    }, {"path": "internal/cpu", "hash": "b5b31e679dd3cfb6dc787f481dccfaa5c191e967"}, {
        "path": "internal/bytealg",
        "hash": "7b0097da94458f900c3f593199b228134e4da300"
    }, {"path": "runtime/internal/sys", "hash": "39c26399487a21130e04214c580d469c7f8d5f57"}, {
        "path": "runtime",
        "hash": "2d657268b5e6b21ea671e5a50437f6511122a671"
    }, {"path": "errors", "hash": "92509deefa51b9d6ded6a744ffbdb76da00c2349"}, {
        "path": "internal/race",
        "hash": "2a5898f657203ab95a30f72ace9132934cd44b50"
    }, {"path": "sync/atomic", "hash": "03c2b4d4a1a61c0cde76907772e28510c845ecad"}, {
        "path": "sync",
        "hash": "797e7c045f341736adf97257cd9f4199de701c47"
    }, {"path": "io", "hash": "347b6e6e815e50ab600fa50dc9980914124e3af8"}, {
        "path": "unicode",
        "hash": "703b1e6f7af8a8f4e0dd7bbf31cae164dd667973"
    }, {"path": "unicode/utf8", "hash": "d7beb80cd54f3b382af915b4ba61a9893b4a7fb7"}, {
        "path": "bytes",
        "hash": "b23aa2c2bcfc597083ad49cfe19ffa8035aa6032"
    }, {"path": "bufio", "hash": "55bc5eecd73afca6bb141dc8984048b112a35e80"}, {
        "path": "image/color",
        "hash": "efde565dd804d0a0d5a0020730306a4721c9b182"
    }, {"path": "math", "hash": "3c34e5698995754d3e057a9553c4a47cb0b7e03e"}, {
        "path": "math/bits",
        "hash": "17081cf6e74fd306210d9ed4d4ac28466a64581e"
    }, {"path": "strconv", "hash": "6a4accd68bbaa3cbffbeabaf583b13067fb5b6c0"}, {
        "path": "image",
        "hash": "6a17635cf81d5c1e4adc192472197cbf4703ae67"
    }, {
        "path": "strings",
        "hash": "da36c9388403366cc805755f7ac03911aa9f4273"
    }, {
        "path": "github.com/gopherjs/gopherjs/nosync",
        "hash": "8dc6b1af74253ec326cf1ec604c03b6ab2e8c606"
    }, {"path": "syscall", "hash": "cb166b41d21830de41d2f5f699ad10421e4f9cd4"}, {
        "path": "time",
        "hash": "fd7c9ec7695cf6a1b803e329df7fa8cba1c9e83a"
    }, {
        "path": "honnef.co/go/js/dom",
        "hash": "d1d2f774adbcb22cc4fbf3c6388ce8e5e594a8e7"
    }, {
        "path": "github.com/MJKWoolnough/gopherjs/files",
        "hash": "2271ad3f0370d6bf043077272332b502e7439284"
    }, {"path": "internal/poll", "hash": "c4fc8031fdeb4b930310a780b480d7fb6bef57c5"}, {
        "path": "internal/syscall/unix",
        "hash": "08bfb320e811001dbe2738c98f394fe158a94151"
    }, {"path": "internal/testlog", "hash": "cf3431b2ce69cd1e27d2a87a47eae6b89f594635"}, {
        "path": "os",
        "hash": "8540fe8b944db5c0f329305a6710a4170125e2a0"
    }, {"path": "reflect", "hash": "f34e3a5ad9c1c23fbc5a0b345c7d1412c6459644"}, {
        "path": "sort",
        "hash": "1a8eddcbf7f1a1be551552ff46a33dde64d9996c"
    }, {
        "path": "path/filepath",
        "hash": "9d1650eca84e8c2caff7db568c641f1bb5a3d03a"
    }, {"path": "github.com/dave/dropper", "hash": "c218d79cbd2f4f6eb864ffca3be36d91605c3981"}, {
        "path": "fmt",
        "hash": "983288ecf3d6bdc5e9424fafc95253c6255ed019"
    }, {
        "path": "github.com/dave/flux",
        "hash": "2edb063e4dbf644cd6e4b0f445947c08fc0220a4"
    }, {
        "path": "github.com/dave/play/models",
        "hash": "87d13fc300aba2c18f694397fdc025a8b200070e"
    }, {
        "path": "github.com/dave/services",
        "hash": "86751bca2f63c92c676742ba4e72befbdc55e0e0"
    }, {
        "path": "github.com/dave/play/actions",
        "hash": "b5b909c94ed273363c587452a95a3ed1dce768c9"
    }, {"path": "compress/flate", "hash": "2326c4b9a5105340a18ef14cec8636f099286328"}, {
        "path": "encoding/binary",
        "hash": "097a1d1e53f75519cbfa10092701a58c9d3f2ae8"
    }, {"path": "hash", "hash": "b2000854f603111744f8bfcb2146786e1c4b6da9"}, {
        "path": "hash/crc32",
        "hash": "56532d3e379da5948282d8041eb55b95b422b767"
    }, {"path": "io/ioutil", "hash": "ef0e4416ca42f4e9b864c5cad22f2565f29a508a"}, {
        "path": "path",
        "hash": "8d8ebdbad1048a04fa5415f967f40cea588f98bc"
    }, {"path": "archive/zip", "hash": "d2be0b4a1b2000d6bcd8b67e4b51279d19d9a210"}, {
        "path": "context",
        "hash": "83ac71c63e6f9cb792b91a7f7b6888f996dd6097"
    }, {"path": "encoding", "hash": "990819c679436759610b5d7846c3e21cb80ca608"}, {
        "path": "encoding/gob",
        "hash": "556ac82a396e16dac6605eedf85638a30de3fe46"
    }, {"path": "encoding/base64", "hash": "7500dc96e009bd0727dc3d2ba9879d459a9abce8"}, {
        "path": "unicode/utf16",
        "hash": "8c816d2a74a2421c47d75a5d866a7218a7d4c8ce"
    }, {
        "path": "encoding/json",
        "hash": "3d87bdfb11b8bb86883659ca90f848caba6485ee"
    }, {
        "path": "github.com/dave/jsgo/config",
        "hash": "ce29bd11cf0c5087144454099ec7b614096e3055"
    }, {
        "path": "github.com/dave/jsgo/server/servermsg",
        "hash": "6436f8e3eaea2c2ffa28ea63794e2310e126b6ec"
    }, {
        "path": "github.com/dave/services/builder/buildermsg",
        "hash": "d91abb37030cb4963d28874121f70561f0c31ced"
    }, {
        "path": "github.com/dave/services/constor/constormsg",
        "hash": "409bf36a4e96947250f1f60d502fc2aee6bda743"
    }, {
        "path": "github.com/dave/services/deployer/deployermsg",
        "hash": "dce9989d743561f49316113e0c82823d783feb5e"
    }, {
        "path": "github.com/dave/services/getter/gettermsg",
        "hash": "fc12d507c0e19e8140b554387d8d9c13dbd1bdce"
    }, {"path": "math/rand", "hash": "1c7b04835516df24a928fdc07c7a6612e41dd94d"}, {
        "path": "math/big",
        "hash": "7ebbe45ca04bfcb13be67aab0d360262bea0c05b"
    }, {"path": "crypto/rand", "hash": "4399eb66fad2018b75cc52abff7efdba64f3cc3f"}, {
        "path": "crypto",
        "hash": "e8d64fac6dd9eddc09c83f9f7df90a444f46f397"
    }, {"path": "crypto/sha1", "hash": "86f171adbe555dafa0d640279f18e3ec89b5e0d2"}, {
        "path": "container/list",
        "hash": "2fecd34d187d924db2a515a1ca3fcee4886cec57"
    }, {"path": "crypto/internal/subtle", "hash": "33e6f094023a8764159ee792e7ca5c866683a562"}, {
        "path": "crypto/subtle",
        "hash": "e5a41904d471f25595a5b5eb5b92a41b762d93f7"
    }, {"path": "crypto/cipher", "hash": "ce73478ebd84850719128ac939fa0934ceb9100e"}, {
        "path": "crypto/aes",
        "hash": "5ca268832735d186f1cffac74ee35362bcaffb86"
    }, {"path": "crypto/des", "hash": "8cb7cae3b8c1f18cb673729edbbcb61e3565a57d"}, {
        "path": "crypto/elliptic",
        "hash": "9a9011f037b757c798903737bbc46fa56f74dcc5"
    }, {
        "path": "crypto/internal/randutil",
        "hash": "08784aa1adb49faeb6e45bcde4afe44dadd99630"
    }, {"path": "crypto/sha512", "hash": "0092815894b6c526533dff50c60ad5c39efdef26"}, {
        "path": "encoding/asn1",
        "hash": "00a03c571727fb101f2f1d7acff4ee0e57ef9b92"
    }, {"path": "crypto/ecdsa", "hash": "01fdacc54e1c513ebc272aac6e827586bb6802aa"}, {
        "path": "crypto/hmac",
        "hash": "906b9ea14281071f999a6c8b3797ae5bc886e3d6"
    }, {"path": "crypto/md5", "hash": "3b91acacf81ad59fcfb41732a1eec44960cc13ea"}, {
        "path": "crypto/rc4",
        "hash": "386d093dcf7c2a5c9473ef0dc1af753878a36800"
    }, {"path": "crypto/rsa", "hash": "c0063721b4422af455bd5b8f69377cfce4466cb1"}, {
        "path": "crypto/sha256",
        "hash": "c50d3dfdc10e23b3e70b20b22cb07cbc70bd3c48"
    }, {"path": "crypto/dsa", "hash": "7fffc42b72b36f2c7ab8837011475b428c44613c"}, {
        "path": "encoding/hex",
        "hash": "61c898c7f8fcfd4f54439581309142624def204c"
    }, {"path": "crypto/x509/pkix", "hash": "ec8141aaab3fca2dc2820b3afbac9e2fe3c813d2"}, {
        "path": "encoding/pem",
        "hash": "a634ae28c5effde08b72b519d5866adccfe0b91e"
    }, {
        "path": "golang_org/x/crypto/cryptobyte/asn1",
        "hash": "43121b68eb3c23a9ec6422df8036e9125b6b9680"
    }, {
        "path": "golang_org/x/crypto/cryptobyte",
        "hash": "4179f3c91359c4b0e717d78c5eab1ee877fe8396"
    }, {
        "path": "golang_org/x/net/dns/dnsmessage",
        "hash": "92f458d863372af1c7c104ec3df41d8b7ba2c9a7"
    }, {
        "path": "golang_org/x/net/route",
        "hash": "9f790b9a77bd8fe6fc42f009ece7467c80bcc08d"
    }, {
        "path": "internal/nettrace",
        "hash": "f3523768c2158dafb37224733fad761d67fe4016"
    }, {"path": "internal/singleflight", "hash": "780922bf96d7d0cf98a0ed483bd1d0df3f4f6187"}, {
        "path": "net",
        "hash": "8ece59bc23ca06e3182184064457bab873b5441f"
    }, {"path": "net/url", "hash": "2e64f8ca344498fc06153454c23ed3151f449094"}, {
        "path": "os/exec",
        "hash": "4644539373a241545b6f4e2d805c65903e54a73e"
    }, {"path": "os/user", "hash": "48b2b595e76b131d677bdc1c7a83fbf250b2f57c"}, {
        "path": "crypto/x509",
        "hash": "4bc99bef1e1eb2b4da3dcd5c76db806cb7b791c9"
    }, {
        "path": "golang_org/x/crypto/internal/chacha20",
        "hash": "d6bbc6c919cb95967fe6ac5963d920a43f620a6d"
    }, {
        "path": "golang_org/x/crypto/poly1305",
        "hash": "9a9fe85f8d5626d50c60f0b516c14f1c9b60175b"
    }, {
        "path": "golang_org/x/crypto/chacha20poly1305",
        "hash": "8a0b7f1f1b93ca4f6b88f53290d990ee01399db1"
    }, {
        "path": "golang_org/x/crypto/curve25519",
        "hash": "b76fd0cd8bd3495d9269c468536fd47dfb8a3f4d"
    }, {"path": "crypto/tls", "hash": "9cd779f06ee33c6eedf9aaaee974f1794037671f"}, {
        "path": "compress/gzip",
        "hash": "4d712f8567960a36a1f90221e899dc5124aad139"
    }, {"path": "golang_org/x/text/transform", "hash": "bd166ef3f3e1030e00d85673e77152794dee55a4"}, {
        "path": "log",
        "hash": "72a5750d64e0e0ad83f9237d21fc57a7ea031811"
    }, {
        "path": "golang_org/x/text/unicode/bidi",
        "hash": "cb9a8f4466beea358e358d16bcbcee99b7fce039"
    }, {
        "path": "golang_org/x/text/secure/bidirule",
        "hash": "372a59192ea6d9d197387c2bf34517054902e154"
    }, {
        "path": "golang_org/x/text/unicode/norm",
        "hash": "254caad5290e57a90d7f2c99c64baabe262a4eac"
    }, {"path": "golang_org/x/net/idna", "hash": "1beb9074336d0afabd867aaececf35264022723e"}, {
        "path": "net/textproto",
        "hash": "158f9657a78034f5bc5facc5648537cb4b843e35"
    }, {
        "path": "golang_org/x/net/http/httpguts",
        "hash": "e8af8d556f6a34041006661bcf04bb922bc57c14"
    }, {
        "path": "golang_org/x/net/http/httpproxy",
        "hash": "68b5430bbc66f6136b5474b134ab8079823ea6c7"
    }, {"path": "golang_org/x/net/http2/hpack", "hash": "25ceaea3f8868b0546cd03e1eff60bb357f6777d"}, {
        "path": "mime",
        "hash": "20ffdf1d8fad0c28fece7e0ebf87d26c6275f474"
    }, {"path": "mime/quotedprintable", "hash": "f8d1186a1f0420e73077c4812afddfb21ff0d4ae"}, {
        "path": "mime/multipart",
        "hash": "ca213322221f83e7800f22043f9589af431eac62"
    }, {"path": "net/http/httptrace", "hash": "92400c0a4e3334ba56583856cd5e394acf9e9fb2"}, {
        "path": "net/http/internal",
        "hash": "4bc2cbf9fc37ee01f38bd3e136e0f5a2aaf365b4"
    }, {
        "path": "net/http",
        "hash": "3725e4f1ea84250e0bb5bfb2a7f38406a64ec06e"
    }, {
        "path": "github.com/gorilla/websocket",
        "hash": "387b21d74b09ac745bd03a9f302b3deb759ff5db"
    }, {
        "path": "github.com/dave/jsgo/server/play/messages",
        "hash": "3add4296185c87c9edb8f5f3a104fc20bea9e3e0"
    }, {"path": "github.com/dave/locstor", "hash": "db8e4a2f259c0d5c39eb410cd6ee4d078ad110aa"}, {
        "path": "go/token",
        "hash": "3d114a16d30de9f5b2b8d13ffe0e332b2f66f3fc"
    }, {"path": "go/scanner", "hash": "3618e0e81e6f52bb145e0aeba603ccce4574e853"}, {
        "path": "go/ast",
        "hash": "cb6f3958d65d51d2052718657013885c87669d92"
    }, {"path": "regexp/syntax", "hash": "31225f61651f2ad8467e8fb70cfee4e6b78c5f7d"}, {
        "path": "regexp",
        "hash": "2c111de93cdec6b6397aeea0e00a91d8db260109"
    }, {"path": "text/template/parse", "hash": "748d4a3ae86bb6c661c861cd587b6b25cd121098"}, {
        "path": "text/template",
        "hash": "9b1573c3e336665b3b8f6fcc852fb11c422976d9"
    }, {"path": "go/doc", "hash": "2e3a8c61859234e67b4d953a953b177b6875a3cc"}, {
        "path": "go/parser",
        "hash": "e744c4f68488fb88e8e9d4c70aa842ca9bd86c8b"
    }, {
        "path": "go/build",
        "hash": "ef3adbd08bc6b9a28a013c0b4c8ec2b75d021d79"
    }, {
        "path": "github.com/dave/services/includer",
        "hash": "d0424006d52dd5e5b9a9760fcd027b0831cc1f30"
    }, {"path": "container/heap", "hash": "755cd855e3be7261dc30bd349c73d8f128a4252b"}, {
        "path": "go/constant",
        "hash": "fd7cf14ab095917d0b97785087c1d69950a6baab"
    }, {
        "path": "go/types",
        "hash": "4b47e87b14a0400e4f87fc5bd49e65bcc56d3b4a"
    }, {
        "path": "github.com/gopherjs/gopherjs/compiler/astutil",
        "hash": "c60ea8cd6677bd3ac2096681d8f30fcd131ea686"
    }, {
        "path": "github.com/gopherjs/gopherjs/compiler/typesutil",
        "hash": "39ecc46578dbe831231a6ad72cbb240c6e8ae916"
    }, {
        "path": "github.com/gopherjs/gopherjs/compiler/analysis",
        "hash": "a7339e5083801795cecf73c09e51bc8be8382c8b"
    }, {
        "path": "github.com/gopherjs/gopherjs/compiler/filter",
        "hash": "79ea4fdd110c0c889d0c7410e442d09e375ae1ea"
    }, {
        "path": "github.com/gopherjs/gopherjs/compiler/prelude",
        "hash": "fe7884813db35fb721d5b2fc52af9f050cd45c52"
    }, {
        "path": "github.com/neelance/astrewrite",
        "hash": "cea18392423691b471d1c73761130f5f1ba33ac3"
    }, {
        "path": "text/scanner",
        "hash": "51953000202d6bc29674062050e732c97df266ca"
    }, {
        "path": "golang.org/x/tools/go/internal/gcimporter",
        "hash": "b8c0e94d10c0ac6eb853270b6ddedf65385ad321"
    }, {
        "path": "golang.org/x/tools/go/gcexportdata",
        "hash": "e4e97503d2296b829ef25142f665749107de6969"
    }, {
        "path": "golang.org/x/tools/go/types/typeutil",
        "hash": "bbe85e68658f5079752ee5226f33c325ad643886"
    }, {
        "path": "github.com/gopherjs/gopherjs/compiler",
        "hash": "6dffa4678ca176dd34e116097384d59205828673"
    }, {
        "path": "github.com/dave/play/stores/builderjs",
        "hash": "71e4de36e89b453b471ff1560757b45da1c5ed55"
    }, {
        "path": "github.com/flimzy/jsblob",
        "hash": "f1a6c25e54272cdeef5ed6355fc5bfef84c06daf"
    }, {
        "path": "github.com/dave/saver",
        "hash": "de764543d0b3f2828b1664c30df7c58c278784a1"
    }, {
        "path": "github.com/gopherjs/websocket/websocketjs",
        "hash": "027fa32491e990c89b60cd30961a99d4b2552b60"
    }, {"path": "text/tabwriter", "hash": "1b0b9e06e9c3ff8caa632c9b6ab7b41fbee885d9"}, {
        "path": "go/printer",
        "hash": "c898930f42588d3dbb1405b0f562396740874e18"
    }, {
        "path": "go/format",
        "hash": "61064bd49fdd97af4412fae5b07fd99e01f4597d"
    }, {
        "path": "github.com/dave/play/stores",
        "hash": "5fb5893c5f97dc8ad58faba05943d141ee86230a"
    }, {
        "path": "github.com/dave/splitter",
        "hash": "111d778a9b1f94db282ec17a3c7bbc2d9d5eff8e"
    }, {
        "path": "github.com/gopherjs/vecty",
        "hash": "ca3f705946d0858e85be0c65be8a151e492a2e23"
    }, {
        "path": "github.com/gopherjs/vecty/elem",
        "hash": "188c3bcbf5d34cd352e01e55809f94366f510353"
    }, {
        "path": "github.com/gopherjs/vecty/event",
        "hash": "46a837c760701008451e13140d8f081b22f174ce"
    }, {
        "path": "github.com/gopherjs/vecty/prop",
        "hash": "a68eec37007fb08bebd74aeed6198b6b5d88b6e4"
    }, {
        "path": "github.com/russross/blackfriday",
        "hash": "dc8dba3b0d040fa30f1aed287e3e07dd107c7eaf"
    }, {
        "path": "github.com/tulir/gopher-ace",
        "hash": "638e539927ebc1d1e8c3871f49755c765578315b"
    }, {
        "path": "github.com/dave/play/views",
        "hash": "4bfc03df8d42a1a2a597680414eca5cea20e8ae2"
    }, {
        "path": "github.com/vincent-petithory/dataurl",
        "hash": "010a6483c9a82c86dd655629b237d47f9ba0e1c0"
    }, {
        "path": "github.com/dave/play",
        "hash": "ff283e9e980552b0edaca99403b81107859b42c9"
    }], o = (document.getElementById("log"), function () {
        n++, window.jsgoProgress && window.jsgoProgress(n, t), n == t && function () {
            for (var n = 0; n < e.length; n++) $load[e[n].path]();
            $mainPkg = $packages["github.com/dave/play"], $synthesizeMethods(), $packages.runtime.$init(), $go($mainPkg.$init, []), $flushConsole()
        }()
    }), a = function (n) {
        t++;
        var e = document.createElement("script");
        e.src = n, e.onload = o, e.onreadystatechange = o, document.head.appendChild(e)
    }, s = 0; s < e.length; s++) a("https://pkg.jsgo.io/" + e[s].path + "." + e[s].hash + ".js")
}();