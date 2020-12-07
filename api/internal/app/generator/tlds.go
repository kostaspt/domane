package generator

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

// Note: This file is getting autogenerated by update_tld.sh. Check what is happening there before updating.

func allExtensions() (extensions map[string]string) {
	src := `{"aaa":"aaa","aarp":"aarp","abarth":"abarth","abb":"abb","abbott":"abbott","abbvie":"abbvie","abc":"abc","able":"able","abogado":"abogado","abudhabi":"abudhabi","ac":"ac","academy":"academy","accenture":"accenture","accountant":"accountant","accountants":"accountants","aco":"aco","active":"active","actor":"actor","ad":"ad","adac":"adac","ads":"ads","adult":"adult","ae":"ae","aeg":"aeg","aero":"aero","aetna":"aetna","af":"af","afamilycompany":"afamilycompany","afl":"afl","africa":"africa","ag":"ag","agakhan":"agakhan","agency":"agency","ai":"ai","aig":"aig","aigo":"aigo","airbus":"airbus","airforce":"airforce","airtel":"airtel","akdn":"akdn","al":"al","alfaromeo":"alfaromeo","alibaba":"alibaba","alipay":"alipay","allfinanz":"allfinanz","allstate":"allstate","ally":"ally","alsace":"alsace","alstom":"alstom","am":"am","americanexpress":"americanexpress","americanfamily":"americanfamily","amex":"amex","amfam":"amfam","amica":"amica","amsterdam":"amsterdam","analytics":"analytics","android":"android","anquan":"anquan","anz":"anz","ao":"ao","aol":"aol","apartments":"apartments","app":"app","apple":"apple","aq":"aq","aquarelle":"aquarelle","ar":"ar","arab":"arab","aramco":"aramco","archi":"archi","army":"army","arpa":"arpa","art":"art","arte":"arte","as":"as","asda":"asda","asia":"asia","associates":"associates","at":"at","athleta":"athleta","attorney":"attorney","au":"au","auction":"auction","audi":"audi","audible":"audible","audio":"audio","auspost":"auspost","author":"author","auto":"auto","autos":"autos","avianca":"avianca","aw":"aw","aws":"aws","ax":"ax","axa":"axa","az":"az","azure":"azure","ba":"ba","baby":"baby","baidu":"baidu","banamex":"banamex","bananarepublic":"bananarepublic","band":"band","bank":"bank","bar":"bar","barcelona":"barcelona","barclaycard":"barclaycard","barclays":"barclays","barefoot":"barefoot","bargains":"bargains","baseball":"baseball","basketball":"basketball","bauhaus":"bauhaus","bayern":"bayern","bb":"bb","bbc":"bbc","bbt":"bbt","bbva":"bbva","bcg":"bcg","bcn":"bcn","bd":"bd","be":"be","beats":"beats","beauty":"beauty","beer":"beer","bentley":"bentley","berlin":"berlin","best":"best","bestbuy":"bestbuy","bet":"bet","bf":"bf","bg":"bg","bh":"bh","bharti":"bharti","bi":"bi","bible":"bible","bid":"bid","bike":"bike","bing":"bing","bingo":"bingo","bio":"bio","biz":"biz","bj":"bj","black":"black","blackfriday":"blackfriday","blanco":"blanco","blockbuster":"blockbuster","blog":"blog","bloomberg":"bloomberg","blue":"blue","bm":"bm","bms":"bms","bmw":"bmw","bn":"bn","bnl":"bnl","bnpparibas":"bnpparibas","bo":"bo","boats":"boats","boehringer":"boehringer","bofa":"bofa","bom":"bom","bond":"bond","boo":"boo","book":"book","booking":"booking","bosch":"bosch","bostik":"bostik","boston":"boston","bot":"bot","boutique":"boutique","box":"box","br":"br","bradesco":"bradesco","bridgestone":"bridgestone","broadway":"broadway","broker":"broker","brother":"brother","brussels":"brussels","bs":"bs","bt":"bt","budapest":"budapest","bugatti":"bugatti","build":"build","builders":"builders","business":"business","buy":"buy","buzz":"buzz","bv":"bv","bw":"bw","by":"by","bz":"bz","bzh":"bzh","ca":"ca","cab":"cab","cafe":"cafe","cal":"cal","call":"call","calvinklein":"calvinklein","cam":"cam","camera":"camera","camp":"camp","cancerresearch":"cancerresearch","canon":"canon","capetown":"capetown","capital":"capital","capitalone":"capitalone","car":"car","caravan":"caravan","cards":"cards","care":"care","career":"career","careers":"careers","cars":"cars","cartier":"cartier","casa":"casa","case":"case","caseih":"caseih","cash":"cash","casino":"casino","cat":"cat","catering":"catering","catholic":"catholic","cba":"cba","cbn":"cbn","cbre":"cbre","cbs":"cbs","cc":"cc","cd":"cd","ceb":"ceb","center":"center","ceo":"ceo","cern":"cern","cf":"cf","cfa":"cfa","cfd":"cfd","cg":"cg","ch":"ch","chanel":"chanel","channel":"channel","chase":"chase","chat":"chat","cheap":"cheap","chintai":"chintai","christmas":"christmas","chrome":"chrome","chrysler":"chrysler","church":"church","ci":"ci","cipriani":"cipriani","circle":"circle","cisco":"cisco","citadel":"citadel","citi":"citi","citic":"citic","city":"city","cityeats":"cityeats","ck":"ck","cl":"cl","claims":"claims","cleaning":"cleaning","click":"click","clinic":"clinic","clinique":"clinique","clothing":"clothing","cloud":"cloud","club":"club","clubmed":"clubmed","cm":"cm","cn":"cn","co":"co","coach":"coach","codes":"codes","coffee":"coffee","college":"college","cologne":"cologne","com":"com","comcast":"comcast","commbank":"commbank","community":"community","company":"company","compare":"compare","computer":"computer","comsec":"comsec","condos":"condos","construction":"construction","consulting":"consulting","contact":"contact","contractors":"contractors","cooking":"cooking","cookingchannel":"cookingchannel","cool":"cool","coop":"coop","corsica":"corsica","country":"country","coupon":"coupon","coupons":"coupons","courses":"courses","cr":"cr","credit":"credit","creditcard":"creditcard","creditunion":"creditunion","cricket":"cricket","crown":"crown","crs":"crs","cruise":"cruise","cruises":"cruises","csc":"csc","cu":"cu","cuisinella":"cuisinella","cv":"cv","cw":"cw","cx":"cx","cy":"cy","cymru":"cymru","cyou":"cyou","cz":"cz","dabur":"dabur","dad":"dad","dance":"dance","data":"data","date":"date","dating":"dating","datsun":"datsun","day":"day","dclk":"dclk","dds":"dds","de":"de","deal":"deal","dealer":"dealer","deals":"deals","degree":"degree","delivery":"delivery","dell":"dell","deloitte":"deloitte","delta":"delta","democrat":"democrat","dental":"dental","dentist":"dentist","desi":"desi","design":"design","dev":"dev","dhl":"dhl","diamonds":"diamonds","diet":"diet","digital":"digital","direct":"direct","directory":"directory","discount":"discount","discover":"discover","dish":"dish","diy":"diy","dj":"dj","dk":"dk","dm":"dm","dnp":"dnp","do":"do","docs":"docs","doctor":"doctor","dodge":"dodge","dog":"dog","doha":"doha","domains":"domains","dot":"dot","download":"download","drive":"drive","dtv":"dtv","dubai":"dubai","duck":"duck","dunlop":"dunlop","duns":"duns","dupont":"dupont","durban":"durban","dvag":"dvag","dvr":"dvr","dz":"dz","earth":"earth","eat":"eat","ec":"ec","eco":"eco","edeka":"edeka","edu":"edu","education":"education","ee":"ee","eg":"eg","email":"email","emerck":"emerck","energy":"energy","engineer":"engineer","engineering":"engineering","enterprises":"enterprises","epost":"epost","epson":"epson","equipment":"equipment","er":"er","ericsson":"ericsson","erni":"erni","es":"es","esq":"esq","estate":"estate","esurance":"esurance","et":"et","etisalat":"etisalat","eu":"eu","eurovision":"eurovision","eus":"eus","events":"events","everbank":"everbank","exchange":"exchange","expert":"expert","exposed":"exposed","express":"express","extraspace":"extraspace","fage":"fage","fail":"fail","fairwinds":"fairwinds","faith":"faith","family":"family","fan":"fan","fans":"fans","farm":"farm","farmers":"farmers","fashion":"fashion","fast":"fast","fedex":"fedex","feedback":"feedback","ferrari":"ferrari","ferrero":"ferrero","fi":"fi","fiat":"fiat","fidelity":"fidelity","fido":"fido","film":"film","final":"final","finance":"finance","financial":"financial","fire":"fire","firestone":"firestone","firmdale":"firmdale","fish":"fish","fishing":"fishing","fit":"fit","fitness":"fitness","fj":"fj","fk":"fk","flickr":"flickr","flights":"flights","flir":"flir","florist":"florist","flowers":"flowers","fly":"fly","fm":"fm","fo":"fo","foo":"foo","food":"food","foodnetwork":"foodnetwork","football":"football","ford":"ford","forex":"forex","forsale":"forsale","forum":"forum","foundation":"foundation","fox":"fox","fr":"fr","free":"free","fresenius":"fresenius","frl":"frl","frogans":"frogans","frontdoor":"frontdoor","frontier":"frontier","ftr":"ftr","fujitsu":"fujitsu","fujixerox":"fujixerox","fun":"fun","fund":"fund","furniture":"furniture","futbol":"futbol","fyi":"fyi","ga":"ga","gal":"gal","gallery":"gallery","gallo":"gallo","gallup":"gallup","game":"game","games":"games","gap":"gap","garden":"garden","gb":"gb","gbiz":"gbiz","gd":"gd","gdn":"gdn","ge":"ge","gea":"gea","gent":"gent","genting":"genting","george":"george","gf":"gf","gg":"gg","ggee":"ggee","gh":"gh","gi":"gi","gift":"gift","gifts":"gifts","gives":"gives","giving":"giving","gl":"gl","glade":"glade","glass":"glass","gle":"gle","global":"global","globo":"globo","gm":"gm","gmail":"gmail","gmbh":"gmbh","gmo":"gmo","gmx":"gmx","gn":"gn","godaddy":"godaddy","gold":"gold","goldpoint":"goldpoint","golf":"golf","goo":"goo","goodhands":"goodhands","goodyear":"goodyear","goog":"goog","google":"google","gop":"gop","got":"got","gov":"gov","gp":"gp","gq":"gq","gr":"gr","grainger":"grainger","graphics":"graphics","gratis":"gratis","green":"green","gripe":"gripe","grocery":"grocery","group":"group","gs":"gs","gt":"gt","gu":"gu","guardian":"guardian","gucci":"gucci","guge":"guge","guide":"guide","guitars":"guitars","guru":"guru","gw":"gw","gy":"gy","hair":"hair","hamburg":"hamburg","hangout":"hangout","haus":"haus","hbo":"hbo","hdfc":"hdfc","hdfcbank":"hdfcbank","health":"health","healthcare":"healthcare","help":"help","helsinki":"helsinki","here":"here","hermes":"hermes","hgtv":"hgtv","hiphop":"hiphop","hisamitsu":"hisamitsu","hitachi":"hitachi","hiv":"hiv","hk":"hk","hkt":"hkt","hm":"hm","hn":"hn","hockey":"hockey","holdings":"holdings","holiday":"holiday","homedepot":"homedepot","homegoods":"homegoods","homes":"homes","homesense":"homesense","honda":"honda","honeywell":"honeywell","horse":"horse","hospital":"hospital","host":"host","hosting":"hosting","hot":"hot","hoteles":"hoteles","hotels":"hotels","hotmail":"hotmail","house":"house","how":"how","hr":"hr","hsbc":"hsbc","ht":"ht","hu":"hu","hughes":"hughes","hyatt":"hyatt","hyundai":"hyundai","ibm":"ibm","icbc":"icbc","ice":"ice","icu":"icu","id":"id","ie":"ie","ieee":"ieee","ifm":"ifm","ikano":"ikano","il":"il","im":"im","imamat":"imamat","imdb":"imdb","immo":"immo","immobilien":"immobilien","in":"in","industries":"industries","infiniti":"infiniti","info":"info","ing":"ing","ink":"ink","institute":"institute","insurance":"insurance","insure":"insure","int":"int","intel":"intel","international":"international","intuit":"intuit","investments":"investments","io":"io","ipiranga":"ipiranga","iq":"iq","ir":"ir","irish":"irish","is":"is","iselect":"iselect","ismaili":"ismaili","ist":"ist","istanbul":"istanbul","it":"it","itau":"itau","itv":"itv","iveco":"iveco","iwc":"iwc","jaguar":"jaguar","java":"java","jcb":"jcb","jcp":"jcp","je":"je","jeep":"jeep","jetzt":"jetzt","jewelry":"jewelry","jio":"jio","jlc":"jlc","jll":"jll","jm":"jm","jmp":"jmp","jnj":"jnj","jo":"jo","jobs":"jobs","joburg":"joburg","jot":"jot","joy":"joy","jp":"jp","jpmorgan":"jpmorgan","jprs":"jprs","juegos":"juegos","juniper":"juniper","kaufen":"kaufen","kddi":"kddi","ke":"ke","kerryhotels":"kerryhotels","kerrylogistics":"kerrylogistics","kerryproperties":"kerryproperties","kfh":"kfh","kg":"kg","kh":"kh","ki":"ki","kia":"kia","kim":"kim","kinder":"kinder","kindle":"kindle","kitchen":"kitchen","kiwi":"kiwi","km":"km","kn":"kn","koeln":"koeln","komatsu":"komatsu","kosher":"kosher","kp":"kp","kpmg":"kpmg","kpn":"kpn","kr":"kr","krd":"krd","kred":"kred","kuokgroup":"kuokgroup","kw":"kw","ky":"ky","kyoto":"kyoto","kz":"kz","la":"la","lacaixa":"lacaixa","ladbrokes":"ladbrokes","lamborghini":"lamborghini","lamer":"lamer","lancaster":"lancaster","lancia":"lancia","lancome":"lancome","land":"land","landrover":"landrover","lanxess":"lanxess","lasalle":"lasalle","lat":"lat","latino":"latino","latrobe":"latrobe","law":"law","lawyer":"lawyer","lb":"lb","lc":"lc","lds":"lds","lease":"lease","leclerc":"leclerc","lefrak":"lefrak","legal":"legal","lego":"lego","lexus":"lexus","lgbt":"lgbt","li":"li","liaison":"liaison","lidl":"lidl","life":"life","lifeinsurance":"lifeinsurance","lifestyle":"lifestyle","lighting":"lighting","like":"like","lilly":"lilly","limited":"limited","limo":"limo","lincoln":"lincoln","linde":"linde","link":"link","lipsy":"lipsy","live":"live","living":"living","lixil":"lixil","lk":"lk","llc":"llc","loan":"loan","loans":"loans","locker":"locker","locus":"locus","loft":"loft","lol":"lol","london":"london","lotte":"lotte","lotto":"lotto","love":"love","lpl":"lpl","lplfinancial":"lplfinancial","lr":"lr","ls":"ls","lt":"lt","ltd":"ltd","ltda":"ltda","lu":"lu","lundbeck":"lundbeck","lupin":"lupin","luxe":"luxe","luxury":"luxury","lv":"lv","ly":"ly","ma":"ma","macys":"macys","madrid":"madrid","maif":"maif","maison":"maison","makeup":"makeup","man":"man","management":"management","mango":"mango","map":"map","market":"market","marketing":"marketing","markets":"markets","marriott":"marriott","marshalls":"marshalls","maserati":"maserati","mattel":"mattel","mba":"mba","mc":"mc","mckinsey":"mckinsey","md":"md","me":"me","med":"med","media":"media","meet":"meet","melbourne":"melbourne","meme":"meme","memorial":"memorial","men":"men","menu":"menu","meo":"meo","merckmsd":"merckmsd","metlife":"metlife","mg":"mg","mh":"mh","miami":"miami","microsoft":"microsoft","mil":"mil","mini":"mini","mint":"mint","mit":"mit","mitsubishi":"mitsubishi","mk":"mk","ml":"ml","mlb":"mlb","mls":"mls","mm":"mm","mma":"mma","mn":"mn","mo":"mo","mobi":"mobi","mobile":"mobile","mobily":"mobily","moda":"moda","moe":"moe","moi":"moi","mom":"mom","monash":"monash","money":"money","monster":"monster","mopar":"mopar","mormon":"mormon","mortgage":"mortgage","moscow":"moscow","moto":"moto","motorcycles":"motorcycles","mov":"mov","movie":"movie","movistar":"movistar","mp":"mp","mq":"mq","mr":"mr","ms":"ms","msd":"msd","mt":"mt","mtn":"mtn","mtr":"mtr","mu":"mu","museum":"museum","mutual":"mutual","mv":"mv","mw":"mw","mx":"mx","my":"my","mz":"mz","na":"na","nab":"nab","nadex":"nadex","nagoya":"nagoya","name":"name","nationwide":"nationwide","natura":"natura","navy":"navy","nba":"nba","nc":"nc","ne":"ne","nec":"nec","net":"net","netbank":"netbank","netflix":"netflix","network":"network","neustar":"neustar","new":"new","newholland":"newholland","news":"news","next":"next","nextdirect":"nextdirect","nexus":"nexus","nf":"nf","nfl":"nfl","ng":"ng","ngo":"ngo","nhk":"nhk","ni":"ni","nico":"nico","nike":"nike","nikon":"nikon","ninja":"ninja","nissan":"nissan","nissay":"nissay","nl":"nl","no":"no","nokia":"nokia","northwesternmutual":"northwesternmutual","norton":"norton","now":"now","nowruz":"nowruz","nowtv":"nowtv","np":"np","nr":"nr","nra":"nra","nrw":"nrw","ntt":"ntt","nu":"nu","nyc":"nyc","nz":"nz","obi":"obi","observer":"observer","off":"off","office":"office","okinawa":"okinawa","olayan":"olayan","olayangroup":"olayangroup","oldnavy":"oldnavy","ollo":"ollo","om":"om","omega":"omega","one":"one","ong":"ong","onl":"onl","online":"online","onyourside":"onyourside","ooo":"ooo","open":"open","oracle":"oracle","orange":"orange","org":"org","organic":"organic","origins":"origins","osaka":"osaka","otsuka":"otsuka","ott":"ott","ovh":"ovh","pa":"pa","page":"page","panasonic":"panasonic","panerai":"panerai","paris":"paris","pars":"pars","partners":"partners","parts":"parts","party":"party","passagens":"passagens","pay":"pay","pccw":"pccw","pe":"pe","pet":"pet","pf":"pf","pfizer":"pfizer","pg":"pg","ph":"ph","pharmacy":"pharmacy","phd":"phd","philips":"philips","phone":"phone","photo":"photo","photography":"photography","photos":"photos","physio":"physio","piaget":"piaget","pics":"pics","pictet":"pictet","pictures":"pictures","pid":"pid","pin":"pin","ping":"ping","pink":"pink","pioneer":"pioneer","pizza":"pizza","pk":"pk","pl":"pl","place":"place","play":"play","playstation":"playstation","plumbing":"plumbing","plus":"plus","pm":"pm","pn":"pn","pnc":"pnc","pohl":"pohl","poker":"poker","politie":"politie","porn":"porn","post":"post","pr":"pr","pramerica":"pramerica","praxi":"praxi","press":"press","prime":"prime","pro":"pro","prod":"prod","productions":"productions","prof":"prof","progressive":"progressive","promo":"promo","properties":"properties","property":"property","protection":"protection","pru":"pru","prudential":"prudential","ps":"ps","pt":"pt","pub":"pub","pw":"pw","pwc":"pwc","py":"py","qa":"qa","qpon":"qpon","quebec":"quebec","quest":"quest","qvc":"qvc","racing":"racing","radio":"radio","raid":"raid","re":"re","read":"read","realestate":"realestate","realtor":"realtor","realty":"realty","recipes":"recipes","red":"red","redstone":"redstone","redumbrella":"redumbrella","rehab":"rehab","reise":"reise","reisen":"reisen","reit":"reit","reliance":"reliance","ren":"ren","rent":"rent","rentals":"rentals","repair":"repair","report":"report","republican":"republican","rest":"rest","restaurant":"restaurant","review":"review","reviews":"reviews","rexroth":"rexroth","rich":"rich","richardli":"richardli","ricoh":"ricoh","rightathome":"rightathome","ril":"ril","rio":"rio","rip":"rip","rmit":"rmit","ro":"ro","rocher":"rocher","rocks":"rocks","rodeo":"rodeo","rogers":"rogers","room":"room","rs":"rs","rsvp":"rsvp","ru":"ru","rugby":"rugby","ruhr":"ruhr","run":"run","rw":"rw","rwe":"rwe","ryukyu":"ryukyu","sa":"sa","saarland":"saarland","safe":"safe","safety":"safety","sakura":"sakura","sale":"sale","salon":"salon","samsclub":"samsclub","samsung":"samsung","sandvik":"sandvik","sandvikcoromant":"sandvikcoromant","sanofi":"sanofi","sap":"sap","sapo":"sapo","sarl":"sarl","sas":"sas","save":"save","saxo":"saxo","sb":"sb","sbi":"sbi","sbs":"sbs","sc":"sc","sca":"sca","scb":"scb","schaeffler":"schaeffler","schmidt":"schmidt","scholarships":"scholarships","school":"school","schule":"schule","schwarz":"schwarz","science":"science","scjohnson":"scjohnson","scor":"scor","scot":"scot","sd":"sd","se":"se","search":"search","seat":"seat","secure":"secure","security":"security","seek":"seek","select":"select","sener":"sener","services":"services","ses":"ses","seven":"seven","sew":"sew","sex":"sex","sexy":"sexy","sfr":"sfr","sg":"sg","sh":"sh","shangrila":"shangrila","sharp":"sharp","shaw":"shaw","shell":"shell","shia":"shia","shiksha":"shiksha","shoes":"shoes","shop":"shop","shopping":"shopping","shouji":"shouji","show":"show","showtime":"showtime","shriram":"shriram","si":"si","silk":"silk","sina":"sina","singles":"singles","site":"site","sj":"sj","sk":"sk","ski":"ski","skin":"skin","sky":"sky","skype":"skype","sl":"sl","sling":"sling","sm":"sm","smart":"smart","smile":"smile","sn":"sn","sncf":"sncf","so":"so","soccer":"soccer","social":"social","softbank":"softbank","software":"software","sohu":"sohu","solar":"solar","solutions":"solutions","song":"song","sony":"sony","soy":"soy","space":"space","spiegel":"spiegel","sport":"sport","spot":"spot","spreadbetting":"spreadbetting","sr":"sr","srl":"srl","srt":"srt","st":"st","stada":"stada","staples":"staples","star":"star","starhub":"starhub","statebank":"statebank","statefarm":"statefarm","statoil":"statoil","stc":"stc","stcgroup":"stcgroup","stockholm":"stockholm","storage":"storage","store":"store","stream":"stream","studio":"studio","study":"study","style":"style","su":"su","sucks":"sucks","supplies":"supplies","supply":"supply","support":"support","surf":"surf","surgery":"surgery","suzuki":"suzuki","sv":"sv","swatch":"swatch","swiftcover":"swiftcover","swiss":"swiss","sx":"sx","sy":"sy","sydney":"sydney","symantec":"symantec","systems":"systems","sz":"sz","tab":"tab","taipei":"taipei","talk":"talk","taobao":"taobao","target":"target","tatamotors":"tatamotors","tatar":"tatar","tattoo":"tattoo","tax":"tax","taxi":"taxi","tc":"tc","tci":"tci","td":"td","tdk":"tdk","team":"team","tech":"tech","technology":"technology","tel":"tel","telecity":"telecity","telefonica":"telefonica","temasek":"temasek","tennis":"tennis","teva":"teva","tf":"tf","tg":"tg","th":"th","thd":"thd","theater":"theater","theatre":"theatre","tiaa":"tiaa","tickets":"tickets","tienda":"tienda","tiffany":"tiffany","tips":"tips","tires":"tires","tirol":"tirol","tj":"tj","tjmaxx":"tjmaxx","tjx":"tjx","tk":"tk","tkmaxx":"tkmaxx","tl":"tl","tm":"tm","tmall":"tmall","tn":"tn","to":"to","today":"today","tokyo":"tokyo","tools":"tools","top":"top","toray":"toray","toshiba":"toshiba","total":"total","tours":"tours","town":"town","toyota":"toyota","toys":"toys","tr":"tr","trade":"trade","trading":"trading","training":"training","travel":"travel","travelchannel":"travelchannel","travelers":"travelers","travelersinsurance":"travelersinsurance","trust":"trust","trv":"trv","tt":"tt","tube":"tube","tui":"tui","tunes":"tunes","tushu":"tushu","tv":"tv","tvs":"tvs","tw":"tw","tz":"tz","ua":"ua","ubank":"ubank","ubs":"ubs","uconnect":"uconnect","ug":"ug","uk":"uk","unicom":"unicom","university":"university","uno":"uno","uol":"uol","ups":"ups","us":"us","uy":"uy","uz":"uz","va":"va","vacations":"vacations","vana":"vana","vanguard":"vanguard","vc":"vc","ve":"ve","vegas":"vegas","ventures":"ventures","verisign":"verisign","versicherung":"versicherung","vet":"vet","vg":"vg","vi":"vi","viajes":"viajes","video":"video","vig":"vig","viking":"viking","villas":"villas","vin":"vin","vip":"vip","virgin":"virgin","visa":"visa","vision":"vision","vista":"vista","vistaprint":"vistaprint","viva":"viva","vivo":"vivo","vlaanderen":"vlaanderen","vn":"vn","vodka":"vodka","volkswagen":"volkswagen","volvo":"volvo","vote":"vote","voting":"voting","voto":"voto","voyage":"voyage","vu":"vu","vuelos":"vuelos","wales":"wales","walmart":"walmart","walter":"walter","wang":"wang","wanggou":"wanggou","warman":"warman","watch":"watch","watches":"watches","weather":"weather","weatherchannel":"weatherchannel","webcam":"webcam","weber":"weber","website":"website","wed":"wed","wedding":"wedding","weibo":"weibo","weir":"weir","wf":"wf","whoswho":"whoswho","wien":"wien","wiki":"wiki","williamhill":"williamhill","win":"win","windows":"windows","wine":"wine","winners":"winners","wme":"wme","wolterskluwer":"wolterskluwer","woodside":"woodside","work":"work","works":"works","world":"world","wow":"wow","ws":"ws","wtc":"wtc","wtf":"wtf","xbox":"xbox","xerox":"xerox","xfinity":"xfinity","xihuan":"xihuan","xin":"xin","xn--11b4c3d":"u0915u0949u092e","xn--1ck2e1b":"u30bbu30fcu30eb","xn--1qqw23a":"u4f5bu5c71","xn--2scrj9c":"u0cadu0cbeu0cb0u0ca4","xn--30rr7y":"u6148u5584","xn--3bst00m":"u96c6u56e2","xn--3ds443g":"u5728u7ebf","xn--3e0b707e":"ud55cuad6d","xn--3hcrj9c":"u0b2du0b3eu0b30u0b24","xn--3oq18vl8pn36a":"u5927u4f17u6c7du8f66","xn--3pxu8k":"u70b9u770b","xn--42c2d9a":"u0e04u0e2du0e21","xn--45br5cyl":"u09adu09beu09f0u09a4","xn--45brj9c":"u09adu09beu09b0u09a4","xn--45q11c":"u516bu5366","xn--4gbrim":"u0645u0648u0642u0639","xn--54b7fta0cc":"u09acu09beu0982u09b2u09be","xn--55qw42g":"u516cu76ca","xn--55qx5d":"u516cu53f8","xn--5su34j936bgsg":"u9999u683cu91ccu62c9","xn--5tzm5g":"u7f51u7ad9","xn--6frz82g":"u79fbu52a8","xn--6qq986b3xl":"u6211u7231u4f60","xn--80adxhks":"u043cu043eu0441u043au0432u0430","xn--80ao21a":"u049bu0430u0437","xn--80aqecdr1a":"u043au0430u0442u043eu043bu0438u043a","xn--80asehdb":"u043eu043du043bu0430u0439u043d","xn--80aswg":"u0441u0430u0439u0442","xn--8y0a063a":"u8054u901a","xn--90a3ac":"u0441u0440u0431","xn--90ae":"u0431u0433","xn--90ais":"u0431u0435u043b","xn--9dbq2a":"u05e7u05d5u05dd","xn--9et52u":"u65f6u5c1a","xn--9krt00a":"u5faeu535a","xn--b4w605ferd":"u6de1u9a6cu9521","xn--bck1b9a5dre4c":"u30d5u30a1u30c3u30b7u30e7u30f3","xn--c1avg":"u043eu0440u0433","xn--c2br7g":"u0928u0947u091f","xn--cck2b3b":"u30b9u30c8u30a2","xn--cg4bki":"uc0bcuc131","xn--clchc0ea0b2g2a9gcd":"u0b9au0bbfu0b99u0bcdu0b95u0baau0bcdu0baau0bc2u0bb0u0bcd","xn--czr694b":"u5546u6807","xn--czrs0t":"u5546u5e97","xn--czru2d":"u5546u57ce","xn--d1acj3b":"u0434u0435u0442u0438","xn--d1alf":"u043cu043au0434","xn--e1a4c":"u0435u044e","xn--eckvdtc9d":"u30ddu30a4u30f3u30c8","xn--efvy88h":"u65b0u95fb","xn--estv75g":"u5de5u884c","xn--fct429k":"u5bb6u96fb","xn--fhbei":"u0643u0648u0645","xn--fiq228c5hs":"u4e2du6587u7f51","xn--fiq64b":"u4e2du4fe1","xn--fiqs8s":"u4e2du56fd","xn--fiqz9s":"u4e2du570b","xn--fjq720a":"u5a31u4e50","xn--flw351e":"u8c37u6b4c","xn--fpcrj9c3d":"u0c2du0c3eu0c30u0c24u0c4d","xn--fzc2c9e2c":"u0dbdu0d82u0d9au0dcf","xn--fzys8d69uvgm":"u96fbu8a0au76c8u79d1","xn--g2xx48c":"u8d2du7269","xn--gckr3f0f":"u30afu30e9u30a6u30c9","xn--gecrj9c":"u0aadu0abeu0ab0u0aa4","xn--gk3at1e":"u901au8ca9","xn--h2breg3eve":"u092du093eu0930u0924u092eu094d","xn--h2brj9c":"u092du093eu0930u0924","xn--h2brj9c8c":"u092du093eu0930u094bu0924","xn--hxt814e":"u7f51u5e97","xn--i1b6b1a6a2e":"u0938u0902u0917u0920u0928","xn--imr513n":"u9910u5385","xn--io0a7i":"u7f51u7edc","xn--j1aef":"u043au043eu043c","xn--j1amh":"u0443u043au0440","xn--j6w193g":"u9999u6e2f","xn--jlq61u9w7b":"u8bfau57fau4e9a","xn--jvr189m":"u98dfu54c1","xn--kcrx77d1x4a":"u98deu5229u6d66","xn--kprw13d":"u53f0u6e7e","xn--kpry57d":"u53f0u7063","xn--kpu716f":"u624bu8868","xn--kput3i":"u624bu673a","xn--l1acc":"u043cu043eu043d","xn--lgbbat1ad8j":"u0627u0644u062cu0632u0627u0626u0631","xn--mgb9awbf":"u0639u0645u0627u0646","xn--mgba3a3ejt":"u0627u0631u0627u0645u0643u0648","xn--mgba3a4f16a":"u0627u06ccu0631u0627u0646","xn--mgba7c0bbn0a":"u0627u0644u0639u0644u064au0627u0646","xn--mgbaakc7dvf":"u0627u062au0635u0627u0644u0627u062a","xn--mgbaam7a8h":"u0627u0645u0627u0631u0627u062a","xn--mgbab2bd":"u0628u0627u0632u0627u0631","xn--mgbai9azgqp6j":"u067eu0627u06a9u0633u062au0627u0646","xn--mgbayh7gpa":"u0627u0644u0627u0631u062fu0646","xn--mgbb9fbpob":"u0645u0648u0628u0627u064au0644u064a","xn--mgbbh1a":"u0628u0627u0631u062a","xn--mgbbh1a71e":"u0628u06beu0627u0631u062a","xn--mgbc0a9azcg":"u0627u0644u0645u063au0631u0628","xn--mgbca7dzdo":"u0627u0628u0648u0638u0628u064a","xn--mgberp4a5d4ar":"u0627u0644u0633u0639u0648u062fu064au0629","xn--mgbgu82a":"u0680u0627u0631u062a","xn--mgbi4ecexp":"u0643u0627u062bu0648u0644u064au0643","xn--mgbpl2fh":"u0633u0648u062fu0627u0646","xn--mgbt3dhd":"u0647u0645u0631u0627u0647","xn--mgbtx2b":"u0639u0631u0627u0642","xn--mgbx4cd0ab":"u0645u0644u064au0633u064au0627","xn--mix891f":"u6fb3u9580","xn--mk1bu44c":"ub2f7ucef4","xn--mxtq1m":"u653fu5e9c","xn--ngbc5azd":"u0634u0628u0643u0629","xn--ngbe9e0a":"u0628u064au062au0643","xn--ngbrx":"u0639u0631u0628","xn--node":"u10d2u10d4","xn--nqv7f":"u673au6784","xn--nqv7fs00ema":"u7ec4u7ec7u673au6784","xn--nyqy26a":"u5065u5eb7","xn--o3cw4h":"u0e44u0e17u0e22","xn--ogbpf8fl":"u0633u0648u0631u064au0629","xn--otu796d":"u62dbu8058","xn--p1acf":"u0440u0443u0441","xn--p1ai":"u0440u0444","xn--pbt977c":"u73e0u5b9d","xn--pgbs0dh":"u062au0648u0646u0633","xn--pssy2u":"u5927u62ff","xn--q9jyb4c":"u307fu3093u306a","xn--qcka1pmc":"u30b0u30fcu30b0u30eb","xn--qxam":"u03b5u03bb","xn--rhqv96g":"u4e16u754c","xn--rovu88b":"u66f8u7c4d","xn--rvc1e0am3e":"u0d2du0d3eu0d30u0d24u0d02","xn--s9brj9c":"u0a2du0a3eu0a30u0a24","xn--ses554g":"u7f51u5740","xn--t60b56a":"ub2f7ub137","xn--tckwe":"u30b3u30e0","xn--tiq49xqyj":"u5929u4e3bu6559","xn--unup4y":"u6e38u620f","xn--vermgensberater-ctb":"vermu00f6gensberater","xn--vermgensberatung-pwb":"vermu00f6gensberatung","xn--vhquv":"u4f01u4e1a","xn--vuq861b":"u4fe1u606f","xn--w4r85el8fhu5dnra":"u5609u91ccu5927u9152u5e97","xn--w4rs40l":"u5609u91cc","xn--wgbh1c":"u0645u0635u0631","xn--wgbl6a":"u0642u0637u0631","xn--xhq521b":"u5e7fu4e1c","xn--xkc2al3hye2a":"u0b87u0bb2u0b99u0bcdu0b95u0bc8","xn--xkc2dl3a5ee0h":"u0b87u0ba8u0bcdu0ba4u0bbfu0bafu0bbe","xn--y9a3aq":"u0570u0561u0575","xn--yfro4i67o":"u65b0u52a0u5761","xn--ygbi2ammx":"u0641u0644u0633u0637u064au0646","xn--zfr164b":"u653fu52a1","xperia":"xperia","xxx":"xxx","xyz":"xyz","yachts":"yachts","yahoo":"yahoo","yamaxun":"yamaxun","yandex":"yandex","ye":"ye","yodobashi":"yodobashi","yoga":"yoga","yokohama":"yokohama","you":"you","youtube":"youtube","yt":"yt","yun":"yun","za":"za","zappos":"zappos","zara":"zara","zero":"zero","zip":"zip","zippo":"zippo","zm":"zm","zone":"zone","zuerich":"zuerich","zw":"zw"}`

	err := json.Unmarshal([]byte(src), &extensions)
	if err != nil {
		log.Err(err).Send()
		return
	}
	return extensions
}