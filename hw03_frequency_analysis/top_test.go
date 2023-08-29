package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var text2 = `У лукоморья дуб зеленый,
	Златая цепь на дубе том:
	И днем и ночью кот ученый
	Всё ходит по цепи кругом;
	Идет направо — песнь заводит,
	Налево — сказку говорит.

	Там чудеса: там леший бродит,
	Русалка на ветвях сидит;
	Там на неведомых дорожках
	Следы невиданных зверей;
	Избушка там на курьих ножках
	Стоит без окон, без дверей;
	Там лес и дол видений полны;
	Там о заре прихлынут волны
	На брег песчаный и пустой,
	И тридцать витязей прекрасных;
	Чредой из вод выходят ясных,
	И с ними дядька их морской;
	Там королевич мимоходом
	Пленяет грозного царя;
	Там в облаках перед народом
	Через леса, через моря
	Колдун несет богатыря;
	В темнице там царевна тужит,
	А бурый волк ей верно служит;
	Там ступа с Бабою Ягой
	Идет, бредет сама собой;
	Там царь Кащей над златом чахнет;
	Там русской дух… там Русью пахнет!
	И там я был, и мед я пил;
	У моря видел дуб зеленый;
	Под ним сидел, и кот ученый
	Свои мне сказки говорил.
	Одну я помню: сказку эту
	Поведаю теперь я свету…

	Дела давно минувших дней,
	Преданья старины глубокой.

	В толпе могучих сыновей,
	С друзьями, в гриднице высокой
	Владимир-солнце пировал;
	Меньшую дочь он выдавал
	За князя храброго Руслана
	И мед из тяжкого стакана
	За их здоровье выпивал.
	Не скоро ели предки наши,
	Не скоро двигались кругом
	Ковши, серебряные чаши
	С кипящим пивом и вином.
	Они веселье в сердце лили,
	Шипела пена по краям,
	Их важно чашники носили
	И низко кланялись гостям.

	Слилися речи в шум невнятный;
	Жужжит гостей веселый круг;
	Но вдруг раздался глас приятный
	И звонких гуслей беглый звук;
	Все смолкли, слушают Баяна:
	И славит сладостный певец
	Людмилу-прелесть, и Руслана,
	И Лелем свитый им венец.

	Но, страстью пылкой утомленный,
	Не ест, не пьет Руслан влюбленный;
	На друга милого глядит,
	Вздыхает, сердится, горит
	И, щипля ус от нетерпенья,
	Считает каждые мгновенья.
	В уныньи, с пасмурным челом,
	За шумным, свадебным столом
	Сидят три витязя младые;
	Безмолвны, за ковшом пустым,
	Забыли кубки круговые,
	И брашна неприятны им;
	Не слышат вещего Баяна;
	Потупили смущенный взгляд:
	То три соперника Руслана;
	В душе несчастные таят
	Любви и ненависти яд.
	Один — Рогдай, воитель смелый,
	Мечом раздвинувший пределы
	Богатых киевских полей;
	Другой — Фарлаф, крикун надменный,
	В пирах никем не побежденный,
	Но воин скромный средь мечей;
	Последний, полный страстной думы,
	Младой хазарский хан Ратмир:
	Все трое бледны и угрюмы,
	И пир веселый им не в пир.

	Вот кончен он; встают рядами,
	Смешались шумными толпами,
	И все глядят на молодых:
	Невеста очи опустила,
	Как будто сердцем приуныла,
	И светел радостный жених.
	Но тень объемлет всю природу,
	Уж близко к полночи глухой;
	Бояре, задремав от меду,
	С поклоном убрались домой.
	Жених в восторге, в упоенье:
	Ласкает он в воображенье
	Стыдливой девы красоту;
	Но с тайным, грустным умиленьем
	Великий князь благословеньем
	Дарует юную чету.

	И вот невесту молодую
	Ведут на брачную постель;
	Огни погасли… и ночную
	Лампаду зажигает Лель.
	Свершились милые надежды,
	Любви готовятся дары;
	Падут ревнивые одежды
	На цареградские ковры…
	Вы слышите ль влюбленный шепот,
	И поцелуев сладкий звук,
	И прерывающийся ропот
	Последней робости?.. Супруг
	Восторги чувствует заране;
	И вот они настали… Вдруг
	Гром грянул, свет блеснул в тумане,
	Лампада гаснет, дым бежит,
	Кругом всё смерклось, всё дрожит,
	И замерла душа в Руслане…
	Всё смолкло. В грозной тишине
	Раздался дважды голос странный,
	И кто-то в дымной глубине
	Взвился чернее мглы туманной…
	И снова терем пуст и тих;
	Встает испуганный жених,
	С лица катится пот остылый;
	Трепеща, хладною рукой
	Он вопрошает мрак немой…
	О горе: нет подруги милой!
	Хватает воздух он пустой;
	Людмилы нет во тьме густой,
	Похищена безвестной силой.

	Ах, если мученик любви
	Страдает страстью безнадежно;
	Хоть грустно жить, друзья мои,
	Однако жить еще возможно.
	Но после долгих, долгих лет
	Обнять влюбленную подругу,
	Желаний, слез, тоски предмет,
	И вдруг минутную супругу
	Навек утратить… о друзья,
	Конечно лучше б умер я!

	Однако жив Руслан несчастный.
	Но что сказал великий князь?
	Сраженный вдруг молвой ужасной,
	На зятя гневом распалясь,
	Его и двор он созывает:
	«Где, где Людмила?» — вопрошает
	С ужасным, пламенным челом.
	Руслан не слышит. «Дети, други!
	Я помню прежние заслуги:
	О, сжальтесь вы над стариком!
	Скажите, кто из вас согласен
	Скакать за дочерью моей?
	Чей подвиг будет не напрасен,
	Тому — терзайся, плачь, злодей!
	Не мог сберечь жены своей! —
	Тому я дам ее в супруги
	С полцарством прадедов моих.
	Кто ж вызовется, дети, други?..»
	«Я!» — молвил горестный жених.
	«Я! я!» — воскликнули с Рогдаем
	Фарлаф и радостный Ратмир:
	«Сейчас коней своих седлаем;
	Мы рады весь изъездить мир.
	Отец наш, не продлим разлуки;
	Не бойся: едем за княжной».
	И с благодарностью немой
	В слезах к ним простирает руки
	Старик, измученный тоской.

	Все четверо выходят вместе;
	Руслан уныньем как убит;
	Мысль о потерянной невесте
	Его терзает и мертвит.
	Садятся на коней ретивых;
	Вдоль берегов Днепра счастливых
	Летят в клубящейся пыли;
	Уже скрываются вдали;
	Уж всадников не видно боле…
	Но долго всё еще глядит
	Великий князь в пустое поле
	И думой им вослед летит.

	Руслан томился молчаливо,
	И смысл и память потеряв.
	Через плечо глядя спесиво
	И важно подбочась, Фарлаф,
	Надувшись, ехал за Русланом.
	Он говорит: «Насилу я
	На волю вырвался, друзья!
	Ну, скоро ль встречусь с великаном?
	Уж то-то крови будет течь,
	Уж то-то жертв любви ревнивой!
	Повеселись, мой верный меч,
	Повеселись, мой конь ретивый!»

	Хазарский хан, в уме своем
	Уже Людмилу обнимая,
	Едва не пляшет над седлом;
	В нем кровь играет молодая,
	Огня надежды полон взор:
	То скачет он во весь опор,
	То дразнит бегуна лихого,
	Кружит, подъемлет на дыбы
	Иль дерзко мчит на холмы снова.

	Рогдай угрюм, молчит — ни слова…
	Страшась неведомой судьбы
	И мучась ревностью напрасной,
	Всех больше беспокоен он,
	И часто взор его ужасный
	На князя мрачно устремлен.

	Соперники одной дорогой
	Все вместе едут целый день.
	Днепра стал темен брег отлогий;
	С востока льется ночи тень;
	Туманы над Днепром глубоким;
	Пора коням их отдохнуть.
	Вот под горой путем широким
	Широкий пересекся путь.
	«Разъедемся, пора! — сказали, —
	Безвестной вверимся судьбе».
	И каждый конь, не чуя стали,
	По воле путь избрал себе.

	Что делаешь, Руслан несчастный,
	Один в пустынной тишине?
	Людмилу, свадьбы день ужасный,
	Всё, мнится, видел ты во сне.
	На брови медный шлем надвинув,
	Из мощных рук узду покинув,
	Ты шагом едешь меж полей,
	И медленно в душе твоей
	Надежда гибнет, гаснет вера.

	Но вдруг пред витязем пещера;
	В пещере свет. Он прямо к ней
	Идет под дремлющие своды,
	Ровесники самой природы.
	Вошел с уныньем: что же зрит?
	В пещере старец; ясный вид,
	Спокойный взор, брада седая;
	Лампада перед ним горит;
	За древней книгой он сидит,
	Ее внимательно читая.
	«Добро пожаловать, мой сын! —
	Сказал с улыбкой он Руслану. —
	Уж двадцать лет я здесь один
	Во мраке старой жизни вяну;
	Но наконец дождался дня,
	Давно предвиденного мною.
	Мы вместе сведены судьбою;
	Садись и выслушай меня.
	Руслан, лишился ты Людмилы;
	Твой твердый дух теряет силы;
	Но зла промчится быстрый миг:
	На время рок тебя постиг.
	С надеждой, верою веселой
	Иди на всё, не унывай;
	Вперед! мечом и грудью смелой
	Свой путь на полночь пробивай.

	Узнай, Руслан: твой оскорбитель
	Волшебник страшный Черномор,
	Красавиц давний похититель,
	Полнощных обладатель гор.
	Еще ничей в его обитель
	Не проникал доныне взор;
	Но ты, злых козней истребитель,
	В нее ты вступишь, и злодей
	Погибнет от руки твоей.
	Тебе сказать не должен боле:
	Судьба твоих грядущих дней,
	Мой сын, в твоей отныне воле».

	Наш витязь старцу пал к ногам
	И в радости лобзает руку.
	Светлеет мир его очам,
	И сердце позабыло муку.
	Вновь ожил он; и вдруг опять
	На вспыхнувшем лице кручина…
	«Ясна тоски твоей причина;
	Но грусть не трудно разогнать, —
	Сказал старик, — тебе ужасна
	Любовь седого колдуна;
	Спокойся, знай: она напрасна
	И юной деве не страшна.
	Он звезды сводит с небосклона,
	Он свистнет — задрожит луна;
	Но против времени закона
	Его наука не сильна.
	Ревнивый, трепетный хранитель
	Замков безжалостных дверей,
	Он только немощный мучитель
	Прелестной пленницы своей.
	Вокруг нее он молча бродит,
	Клянет жестокий жребий свой…
	Но, добрый витязь, день проходит,
	А нужен для тебя покой».

	Руслан на мягкий мох ложится
	Пред умирающим огнем;
	Он ищет позабыться сном,
	Вздыхает, медленно вертится…
	Напрасно! Витязь наконец:
	«Не спится что-то, мой отец!
	Что делать: болен я душою,
	И сон не в сон, как тошно жить.
	Позволь мне сердце освежить
	Твоей беседою святою.
	Прости мне дерзостный вопрос.
	Откройся: кто ты, благодатный,
	Судьбы наперсник непонятный?
	В пустыню кто тебя занес?»

	Вздохнув с улыбкою печальной,
	Старик в ответ: «Любезный сын,
	Уж я забыл отчизны дальной
	Угрюмый край. Природный финн,
	В долинах, нам одним известных,
	Гоняя стадо сел окрестных,
	В беспечной юности я знал
	Одни дремучие дубравы,
	Ручьи, пещеры наших скал
	Да дикой бедности забавы.
	Но жить в отрадной тишине
	Дано не долго было мне.

	Тогда близ нашего селенья,
	Как милый цвет уединенья,
	Жила Наина. Меж подруг
	Она гремела красотою.
	Однажды утренней порою
	Свои стада на темный луг
	Я гнал, волынку надувая;
	Передо мной шумел поток.
	Одна, красавица младая
	На берегу плела венок.
	Меня влекла моя судьбина…
	Ах, витязь, то была Наина!
	Я к ней — и пламень роковой
	За дерзкий взор мне был наградой,
	И я любовь узнал душой
	С ее небесною отрадой,
	С ее мучительной тоской.

	Умчалась года половина;
	Я с трепетом открылся ей,
	Сказал: люблю тебя, Наина.
	Но робкой горести моей
	Наина с гордостью внимала,
	Лишь прелести свои любя,
	И равнодушно отвечала:
	«Пастух, я не люблю тебя!»

	И всё мне дико, мрачно стало:
	Родная куща, тень дубров,
	Веселы игры пастухов —
	Ничто тоски не утешало.
	В уныньи сердце сохло, вяло.
	И наконец задумал я
	Оставить финские поля;
	Морей неверные пучины
	С дружиной братской переплыть
	И бранной славой заслужить
	Вниманье гордое Наины.
	Я вызвал смелых рыбаков
	Искать опасностей и злата.
	Впервые тихий край отцов
	Услышал бранный звук булата
	И шум немирных челноков.
	Я вдаль уплыл, надежды полный,
	С толпой бесстрашных земляков;
	Мы десять лет снега и волны
	Багрили кровию врагов.
	Молва неслась: цари чужбины
	Страшились дерзости моей;
	Их горделивые дружины
	Бежали северных мечей.
	Мы весело, мы грозно бились,
	Делили дани и дары,
	И с побежденными садились
	За дружелюбные пиры.
	Но сердце, полное Наиной,
	Под шумом битвы и пиров,
	Томилось тайною кручиной,
	Искало финских берегов.
	Пора домой, сказал я, други!
	Повесим праздные кольчуги
	Под сенью хижины родной.
	Сказал — и весла зашумели;
	И, страх оставя за собой,
	В залив отчизны дорогой
	Мы с гордой радостью влетели.

	Сбылись давнишние мечты,
	Сбылися пылкие желанья!
	Минута сладкого свиданья,
	И для меня блеснула ты!
	К ногам красавицы надменной
	Принес я меч окровавленный,
	Кораллы, злато и жемчуг;
	Пред нею, страстью упоенный,
	Безмолвным роем окруженный
	Ее завистливых подруг,
	Стоял я пленником послушным;
	Но дева скрылась от меня,
	Примолвя с видом равнодушным:
	«Герой, я не люблю тебя!»

	К чему рассказывать, мой сын,
	Чего пересказать нет силы?
	Ах, и теперь один, один,
	Душой уснув, в дверях могилы,
	Я помню горесть, и порой,
	Как о минувшем мысль родится,
	По бороде моей седой
	Слеза тяжелая катится.

	Но слушай: в родине моей
	Между пустынных рыбарей
	Наука дивная таится.
	Под кровом вечной тишины,
	Среди лесов, в глуши далекой
	Живут седые колдуны;
	К предметам мудрости высокой
	Все мысли их устремлены;
	Всё слышит голос их ужасный,
	Что было и что будет вновь,
	И грозной воле их подвластны
	И гроб и самая любовь.

	И я, любви искатель жадный,
	Решился в грусти безотрадной
	Наину чарами привлечь
	И в гордом сердце девы хладной
	Любовь волшебствами зажечь.
	Спешил в объятия свободы,
	В уединенный мрак лесов;
	И там, в ученьи колдунов,
	Провел невидимые годы.
	Настал давно желанный миг,
	И тайну страшную природы
	Я светлой мыслию постиг:
	Узнал я силу заклинаньям.
	Венец любви, венец желаньям!
	Теперь, Наина, ты моя!
	Победа наша, думал я.
	Но в самом деле победитель
	Был рок, упорный мой гонитель.

	В мечтах надежды молодой,
	В восторге пылкого желанья,
	Творю поспешно заклинанья,
	Зову духов — и в тьме лесной
	Стрела промчалась громовая,
	Волшебный вихорь поднял вой,
	Земля вздрогнула под ногой…
	И вдруг сидит передо мной
	Старушка дряхлая, седая,
	Глазами впалыми сверкая,
	С горбом, с трясучей головой,
	Печальной ветхости картина.
	Ах, витязь, то была Наина!..
	Я ужаснулся и молчал,
	Глазами страшный призрак мерил,
	В сомненье всё еще не верил
	И вдруг заплакал, закричал:
	«Возможно ль! ах, Наина, ты ли!
	Наина, где твоя краса?
	Скажи, ужели небеса
	Тебя так страшно изменили?
	Скажи, давно ль, оставя свет,
	Расстался я с душой и с милой?
	Давно ли?..» «Ровно сорок лет, —
	Был девы роковой ответ, —
	Сегодня семьдесят мне било.
	Что делать, — мне пищит она, —
	Толпою годы пролетели.
	Прошла моя, твоя весна —
	Мы оба постареть успели.
	Но, друг, послушай: не беда
	Неверной младости утрата.
	Конечно, я теперь седа,
	Немножко, может быть, горбата;
	Не то, что в старину была,
	Не так жива, не так мила;
	Зато (прибавила болтунья)
	Открою тайну: я колдунья!»

	И было в самом деле так.
	Немой, недвижный перед нею,
	Я совершенный был дурак
	Со всей премудростью моею.

	Но вот ужасно: колдовство
	Вполне свершилось по несчастью.
	Мое седое божество
	Ко мне пылало новой страстью.
	Скривив улыбкой страшный рот,
	Могильным голосом урод
	Бормочет мне любви признанье.
	Вообрази мое страданье!
	Я трепетал, потупя взор;
	Она сквозь кашель продолжала
	Тяжелый, страстный разговор:
	«Так, сердце я теперь узнала;
	Я вижу, верный друг, оно
	Для нежной страсти рождено;
	Проснулись чувства, я сгораю,
	Томлюсь желаньями любви…
	Приди в объятия мои…
	О милый, милый! умираю…»

	И между тем она, Руслан,
	Мигала томными глазами;
	И между тем за мой кафтан
	Держалась тощими руками;
	И между тем — я обмирал,
	От ужаса зажмуря очи;
	И вдруг терпеть не стало мочи;
	Я с криком вырвался, бежал.
	Она вослед: «О, недостойный!
	Ты возмутил мой век спокойный,
	Невинной девы ясны дни!
	Добился ты любви Наины,
	И презираешь — вот мужчины!
	Изменой дышат все они!
	Увы, сама себя вини;
	Он обольстил меня, несчастный!
	Я отдалась любови страстной…
	Изменник, изверг! о позор!
	Но трепещи, девичий вор!»

	Так мы расстались. С этих пор
	Живу в моем уединенье
	С разочарованной душой;
	И в мире старцу утешенье
	Природа, мудрость и покой.
	Уже зовет меня могила;
	Но чувства прежние свои
	Еще старушка не забыла
	И пламя поздное любви
	С досады в злобу превратила.
	Душою черной зло любя,
	Колдунья старая, конечно,
	Возненавидит и тебя;
	Но горе на земле не вечно».

	Наш витязь с жадностью внимал
	Рассказы старца; ясны очи
	Дремотой легкой не смыкал
	И тихого полета ночи
	В глубокой думе не слыхал.
	Но день блистает лучезарный…
	Со вздохом витязь благодарный
	Объемлет старца-колдуна;
	Душа надеждою полна;
	Выходит вон. Ногами стиснул
	Руслан заржавшего коня,
	В седле оправился, присвистнул.
	«Отец мой, не оставь меня».
	И скачет по пустому лугу.
	Седой мудрец младому другу
	Кричит вослед: «Счастливый путь!
	Прости, люби свою супругу,
	Советов старца не забудь!»`

var text3 = `Call me Ishmael. Some years ago—never mind how long precisely—having little or no money in my purse, 
	and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world.
	It is a way I have of driving off the spleen and regulating the circulation. Whenever I find myself growing grim about the mouth; 
	whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, 
	and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, 
	that it requires a strong moral principle to prevent me from deliberately stepping into the street, 
	and methodically knocking people’s hats off—then, I account it high time to get to sea as soon as I can. 
	This is my substitute for pistol and ball. With a philosophical flourish Cato throws himself upon his sword; 
	I quietly take to the ship. There is nothing surprising in this. If they but knew it, almost all men in their degree, 
	some time or other, cherish very nearly the same feelings towards the ocean with me.
	There now is your insular city of the Manhattoes, belted round by wharves as Indian isles by coral reefs—commerce surrounds it with her surf. 
	Right and left, the streets take you waterward. Its extreme downtown is the battery, where that noble mole is washed by waves, 
	and cooled by breezes, which a few hours previous were out of sight of land. Look at the crowds of water-gazers there.
	Circumambulate the city of a dreamy Sabbath afternoon. Go from Corlears Hook to Coenties Slip, and from thence, by Whitehall, northward. 
	What do you see?—Posted like silent sentinels all around the town, stand thousands upon thousands of mortal men fixed in ocean reveries. 
	Some leaning against the spiles; some seated upon the pier-heads; some looking over the bulwarks of ships from China; some high aloft in the rigging, 
	as if striving to get a still better seaward peep. But these are all landsmen; of week days pent up in lath and plaster—tied to counters, nailed to benches, 
	clinched to desks. How then is this? Are the green fields gone? What do they here?
	But look! here come more crowds, pacing straight for the water, and seemingly bound for a dive. 
	Strange! Nothing will content them but the extremest limit of the land; loitering under the shady lee of yonder warehouses will not suffice. No. 
	They must get just as nigh the water as they possibly can without falling in. And there they stand—miles of them—leagues. Inlanders all, 
	they come from lanes and alleys, streets and avenues—north, east, south, and west. Yet here they all unite. Tell me, 
	does the magnetic virtue of the needles of the compasses of all those ships attract them thither?
	Once more. Say you are in the country; in some high land of lakes. Take almost any path you please, and ten to one it carries you down in a dale,
	and leaves you there by a pool in the stream. There is magic in it. Let the most absent-minded of men be plunged in his deepest reveries—stand that man on his legs, 
	set his feet a-going, and he will infallibly lead you to water, if water there be in all that region. 
	Should you ever be athirst in the great American desert, try this experiment, if your caravan happen to be supplied with a metaphysical professor. 
	Yes, as every one knows, meditation and water are wedded for ever.
	But here is an artist. He desires to paint you the dreamiest, shadiest, quietest, most enchanting bit of romantic landscape in all the valley of the Saco. 
	What is the chief element he employs? There stand his trees, each with a hollow trunk, as if a hermit and a crucifix were within; and here sleeps his meadow, 
	and there sleep his cattle; and up from yonder cottage goes a sleepy smoke. Deep into distant woodlands winds a mazy way, 
	reaching to overlapping spurs of mountains bathed in their hill-side blue. But though the picture lies thus tranced, 
	and though this pine-tree shakes down its sighs like leaves upon this shepherd’s head, yet all were vain, unless the shepherd’s eye were fixed upon the magic stream before him. 
	Go visit the Prairies in June, when for scores on scores of miles you wade knee-deep among Tiger-lilies—what is the one charm wanting?—Water—there is not a drop of water there! 
	Were Niagara but a cataract of sand, would you travel your thousand miles to see it? Why did the poor poet of Tennessee, upon suddenly receiving two handfuls of silver, 
	deliberate whether to buy him a coat, which he sadly needed, or invest his money in a pedestrian trip to Rockaway Beach? 
	Why is almost every robust healthy boy with a robust healthy soul in him, at some time or other crazy to go to sea? Why upon your first voyage as a passenger, 
	did you yourself feel such a mystical vibration, when first told that you and your ship were now out of sight of land? 
	Why did the old Persians hold the sea holy? Why did the Greeks give it a separate deity, and own brother of Jove? Surely all this is not without meaning. 
	And still deeper the meaning of that story of Narcissus, who because he could not grasp the tormenting, mild image he saw in the fountain, 
	plunged into it and was drowned. But that same image, we ourselves see in all rivers and oceans. It is the image of the ungraspable phantom of life; and this is the key to it all.
	Now, when I say that I am in the habit of going to sea whenever I begin to grow hazy about the eyes, and begin to be over conscious of my lungs, 
	I do not mean to have it inferred that I ever go to sea as a passenger. For to go as a passenger you must needs have a purse, 
	and a purse is but a rag unless you have something in it. Besides, passengers get sea-sick—grow quarrelsome—don’t sleep of nights—do not enjoy themselves much, 
	as a general thing;—no, I never go as a passenger; nor, though I am something of a salt, do I ever go to sea as a Commodore, or a Captain, or a Cook. 
	I abandon the glory and distinction of such offices to those who like them. For my part, I abominate all honorable respectable toils, trials, 
	and tribulations of every kind whatsoever. It is quite as much as I can do to take care of myself, without taking care of ships, barques, brigs, 
	schooners, and what not. And as for going as cook,—though I confess there is considerable glory in that, a cook being a sort of officer on ship-board—yet, 
	somehow, I never fancied broiling fowls;—though once broiled, judiciously buttered, and judgmatically salted and peppered, 
	there is no one who will speak more respectfully, not to say reverentially, of a broiled fowl than I will. 
	It is out of the idolatrous dotings of the old Egyptians upon broiled ibis and roasted river horse,
	that you see the mummies of those creatures in their huge bake-houses the pyramids.
	No, when I go to sea, I go as a simple sailor, right before the mast, plumb down into the forecastle, 
	aloft there to the royal mast-head. True, they rather order me about some, and make me jump from spar to spar, like a grasshopper in a May meadow. 
	And at first, this sort of thing is unpleasant enough. It touches one’s sense of honor, particularly if you come of an old established family in the land, 
	the Van Rensselaers, or Randolphs, or Hardicanutes. And more than all, if just previous to putting your hand into the tar-pot, 
	you have been lording it as a country schoolmaster, making the tallest boys stand in awe of you. The transition is a keen one, I assure you,
	from a schoolmaster to a sailor, and requires a strong decoction of Seneca and the Stoics to enable you to grin and bear it. But even this wears off in time.
	What of it, if some old hunks of a sea-captain orders me to get a broom and sweep down the decks? What does that indignity amount to,
	weighed, I mean, in the scales of the New Testament? Do you think the archangel Gabriel thinks anything the less of me, 
	because I promptly and respectfully obey that old hunks in that particular instance? Who ain’t a slave? Tell me that. 
	Well, then, however the old sea-captains may order me about—however they may thump and punch me about, I have the satisfaction of knowing that it is all right; 
	that everybody else is one way or other served in much the same way—either in a physical or metaphysical point of view, that is; 
	and so the universal thump is passed round, and all hands should rub each other’s shoulder-blades, and be content.
	Again, I always go to sea as a sailor, because they make a point of paying me for my trouble, whereas they never pay passengers a single penny that I ever heard of. 
	On the contrary, passengers themselves must pay. And there is all the difference in the world between paying and being paid. 
	The act of paying is perhaps the most uncomfortable infliction that the two orchard thieves entailed upon us. But being paid,—what will compare with it? 
	The urbane activity with which a man receives money is really marvellous, considering that we so earnestly believe money to be the root of all earthly ills, 
	and that on no account can a monied man enter heaven. Ah! how cheerfully we consign ourselves to perdition!
	Finally, I always go to sea as a sailor, because of the wholesome exercise and pure air of the fore-castle deck. For as in this world, 
	head winds are far more prevalent than winds from astern (that is, if you never violate the Pythagorean maxim), 
	so for the most part the Commodore on the quarter-deck gets his atmosphere at second hand from the sailors on the forecastle. 
	He thinks he breathes it first; but not so. In much the same way do the commonalty lead their leaders in many other things, 
	at the same time that the leaders little suspect it. But wherefore it was that after having repeatedly smelt the sea as a merchant sailor, 
	I should now take it into my head to go on a whaling voyage; this the invisible police officer of the Fates, who has the constant surveillance of me, 
	and secretly dogs me, and influences me in some unaccountable way—he can better answer than any one else. And, doubtless, 
	my going on this whaling voyage, formed part of the grand programme of Providence that was drawn up a long time ago. 
	It came in as a sort of brief interlude and solo between more extensive performances. I take it that this part of the bill must have run something like this:
	“Grand Contested Election for the Presidency of the United States. “WHALING VOYAGE BY ONE ISHMAEL. “BLOODY BATTLE IN AFFGHANISTAN.”
	Though I cannot tell why it was exactly that those stage managers, the Fates, put me down for this shabby part of a whaling voyage, 
	when others were set down for magnificent parts in high tragedies, and short and easy parts in genteel comedies, 
	and jolly parts in farces—though I cannot tell why this was exactly; yet, now that I recall all the circumstances, 
	I think I can see a little into the springs and motives which being cunningly presented to me under various disguises, 
	induced me to set about performing the part I did, besides cajoling me into the delusion that it was a choice resulting from my own unbiased freewill and discriminating judgment.
	Chief among these motives was the overwhelming idea of the great whale himself. Such a portentous and mysterious monster roused all my curiosity. 
	Then the wild and distant seas where he rolled his island bulk; the undeliverable, nameless perils of the whale; these,
	with all the attending marvels of a thousand Patagonian sights and sounds, helped to sway me to my wish. With other men, perhaps, 
	such things would not have been inducements; but as for me, I am tormented with an everlasting itch for things remote. 
	I love to sail forbidden seas, and land on barbarous coasts. Not ignoring what is good, I am quick to perceive a horror, 
	and could still be social with it—would they let me—since it is but well to be on friendly terms with all the inmates of the place one lodges in.
	By reason of these things, then, the whaling voyage was welcome; the great flood-gates of the wonder-world swung open,
	and in the wild conceits that swayed me to my purpose, two and two there floated into my inmost soul, endless processions of the whale, 
	and, mid most of them all, one grand hooded phantom, like a snow hill in the air.`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})

	t.Run("positive test 2", func(t *testing.T) {
		expected := []string{
			"И",
			"в",
			"и",
			"не",
			"—",
			"Но",
			"я",
			"В",
			"с",
			"С",
		}

		require.Equal(t, expected, Top10(text2))
	})

	t.Run("positive english test", func(t *testing.T) {
		expected := []string{
			"the",
			"of",
			"a",
			"and",
			"to",
			"in",
			"I",
			"is",
			"that",
			"as",
		}

		require.Equal(t, expected, Top10(text3))
	})
}
