package main

import (
    "fmt"
    "math/rand"
    "reflect"
    "time"
)

type NPC struct {
  name string
  appearance string
  skill string
  talent string
  mannerism string
  flaw string
  ideal string
  race string
}


func gen_name() string {
  prefix := []string{"De","Sin","Kum","Vin","Che","Fre","Roy","Tib","Cen","San","Grab","Gru","Gre","Dir",
    "Vlad","Var","Tok","Tol","Mac","Ada","Sa","Dm","Py","Al","Ar","Jo","Sp","Ti","Gar","Mat","Za","Jer",
    "Att","Bra","Bro","Bri","Re","Ger","Sw","Con","Her","Wi"}

  middle := []string{"il","om","aral","et","er","io","deh","ia","ra","th","ud","ub","row","i","ic","ol"}

  suffix := []string{"an","pi","non","is","ich","us","le","wood","ie","elm","der","osk","oft","obe","eim",
    "ra","on","rey","n","or","el","i","etta","anna","ak","rra","ag","ko","mey","bert"}

  name := prefix[rand.Intn(len(prefix))] + middle[rand.Intn(len(middle))] + suffix[rand.Intn(len(suffix))]
  return name
}


func gen_names(number int){
  for i := 0; i < number; i++{
    fmt.Println(gen_name())
  }
}


func gen_npc() NPC {
    appearance := []string{"lots of jewelery","piercings","flamboyant or colorful clothes",
              "Formal attire", "dirty ratty clothes", "Scar/s", "missing teeth",
              "missing fingers", "unusual eye color or 2 different colors",
              "Tattoos","Birthmark","unusual skin color","Bald","Unusual hair color",
              "nervous twitch", "distinctive nose", "good or bad posture", "beautiful",
              "ugly", "nerdy/ bookworm","long braided blonde hair","very tan","bad posture","very curly hair","allways smiling",
              "allways frowning","wrinkly/old", "child", "fake hand/ wood or metal hook","wear glasses"}

  skills := []string{"strong", "agile/graceful","smart","wise/ spiritual","charasmatic", "hard-ass", "jokster/prankster","weak", "clumsy",
  "dumb", "oblivious", "dull/ boring","night owl"}

  talents := []string{"plays an instrument", "multi langual", "really lucky", "perfect memory",
          "animal lover","great with children","great puzzle solver", "great at one game",
          "great at drawing", "great at painting", "great singer", "great at drinking",
          "master carpenter","expert cook","skilled actor","dancer","knows thieves cant","self taught magic user","pet trainer",
          "ship captain","expert machine tinker","potion maker/herbalist","master thief","master forger","knows the king personally","has a magical item they dont know what it does"}

  mannerisms := []string{"prone to sing, hum or whistle","speaks in rhyme","speak in low or high voice",
             "drunk ","speaks loudly","whispers","uses long words","uses wrong words","fidgets",
             "constantly jokes","expects doomsday","squints(clint eastwood)","spacey (stares off alot)",
             "chewing","pacing","tap fingers","twist hair", "bite fingernails", "makes big hand gestures",
             "never knows the date/ time", "allways excited","act like on crack","a deep thinker", "stretches the truth",
             "clean freak", "horder","stinky","thinks out loud","pregnent/ballsniffer", "allways smoking","cusses alot","smell really good","innaproperate laughing",
             "wants to be an adventurer"}

  flaws := []string{"secret lover ", "prone to falling in love", "enjoys luxury ","Arrogant","envies others things",
        "greedy","prone to rage","has a powerful enemy","phobia","previous scandalous history","secret misdeed",
        "dumbly brave", "lazy", "never surrenders", "married", "in an abusive relationship", "in a relationship", "veryyy single", "single",
        "irrational fear of dumb thing","racist","addicted to drug","sleepwalker","flirty/lustful","abstinate","A bad Bard-barian","experiments on dead bodies",
        "a drow desquised as human","dwarf that wants to be a wizard","a ghost but think they are alive","freaky friday life swapped with a orc warchief","freaky friday life swapped with a Gnome mage",
        "little girl protected by a demon","plants evidence to get people arrested they dont like","secretly a vampire"}

ideals := []string{"dedicated to life goal","protective of family","protective of coworkers","loyal to clan/ employer",
          "deep in love","drawn to a place of interest","protective of a possession","seeking revenge", "very religious"}

race := []string{"human","elf","high elf", "halfling", "gnome", "Tiefling","dwarf","orc","half-elf","half orc"}

  npc_struct := NPC{
    name: gen_name(),
    appearance: appearance[rand.Intn(len(appearance))],
    skill: skills[rand.Intn(len(skills))],
    talent: talents[rand.Intn(len(talents))],
    mannerism: mannerisms[rand.Intn(len(mannerisms))],
    flaw: flaws[rand.Intn(len(flaws))],
    ideal: ideals[rand.Intn(len(ideals))],
    race: race[rand.Intn(len(race))]}

  return npc_struct
}

func gen_npcs(number int){
  for i := 0; i < number; i++{
    v := reflect.ValueOf(gen_npc())
    for ii := 0; ii < v.NumField(); ii++{
      fmt.Println("\t",v.Type().Field(ii).Name ,": \t", v.Field(ii))
    }
    fmt.Println("\n\n")
  }
}

func gen_store(){
  store_types := []string{"magic shop","map shop","armory","blacksmith","food market","general store","clothing store","book store","pet store","jeweler",
        "temple","live stock store", "shipmaker", "carpenter","taxidermist","potion shop","grain mill","bakery","butcher","beekeeper","Church","Casino","Fish/meat Market","art store"}

  store_owner := gen_npc()
  fmt.Println("Store Type: " + store_types[rand.Intn(len(store_types))])
  v := reflect.ValueOf(store_owner)
  for i := 0; i < v.NumField(); i++{
    fmt.Println("\t",v.Type().Field(i).Name ,": \t", v.Field(i))
  }
}


func gen_stores(number int){
  for i := 0; i < number; i++{
    gen_store()
    fmt.Println("\n")
  }
}

func gen_tavern(max int){
  description := []string{"cruel ","shaking ","salty ","juicy ","happy ","green ","red ","grey ",
         "Brass ","gold ","armed ","even ","black ","elven ","silver ","dwarvish ","hellish ","crusty ","silver ","tarnished ", "weathered ",
               "targets ","crumbling ","crimson ","gems ","invisible ","blackened ","charred ","wet ","arrogant "}

  animal := []string{"soldiers ","dragons ","dwarfs ","beholders ","hawks ","hunters ","steeds ","oil lamps ",
         "ravens ","harpy ","manticores ","pixies ","krakens ","giants ","donkeys ",
         "griffon ","rusty ","rats ", "racoons ", "spiders ","banshee ","devils ","lichs ","owlbears ","centaurs "}

  end := []string{"pub","inn","hub","den ","bar ","saloon ","house ","eye ","head ","hideout ","hole ", "watering hole ","lair "}

  tavern_name := "The " + description[rand.Intn(len(description))] + animal[rand.Intn(len(animal))] + end[rand.Intn(len(end))]

  tavern_food := []string{"owlbear ribs","magical carmelized fritters healing potion","Barbecued tiger fish and papaya","Braised beef and pears with ginger","Pork chop & curds","Honey braised boar ribs","Venison and bean stew","Cooked wolf steak",
       "Mushroom stew with corn bread","Acorn soup","Lizard gruel with nutbread","Porridge","Hot beet soup and fresh bread","giant spider stew","Minted pea soup","Elven bread","Braised duck","fried chicken with honey pepper sauce","pasta with red sauce and scallops","roasted roots with wild salad",
       "meat of the day","honey with berries and cornbread","fig feta flatbread","wild salad","potatoe fries","turnip fries","hashbrown and eggs","fried wolf steak and mashed turnips", "Damaran Goulash","Parsnip pie","hellblood biscuits","sweet and spicy meatballs",
       "Garlic Fries","Pot Pie","Drow Mushroom Steak","Fey Wild Salad","Heros Feast Stew","roasted honey glazed turkey","pot meat sandwhich","porkbelly cuts with ripe figs","buckwheat pancakes with mullberry syrup","cooked rabit with ground cherry salsa","underdark mushroom meadly","rye swirl bread"}

  tavern_drinks := []string{"Grog","goblin spit ale","turnip wine","moonshine","dwarven ale","Pulsch brown ale","moon mountian ale","spiced ale","spiced squash stout","kings ale","wrymwizz ale","elderberry spiced wine","lotus leaf wine","stonesulder wine (yellow demon plant)","sweetberry wine",
          "mead","ale","pear sour","grape sparkling wine","moon rum","vodka","scotch","berry brandy","moss mead","honeysuckle mead","coffee stout","fireweed wiskey","mushroom moonshine","greensage cider","abyssal ale","ghost ale(become ghost for hr)","goodale(adv for good chr 1 hr)",
          "sweetwater(cure poison)","rutabega ale","gorgondy ale (see a past vision)","future root ale(see future vision)","beet beer","Firesteap Mountian dew","Cloven Mountian Dew","Black Dragons bite","mushroom tea","ground cherry sour","herbal tomato drink","",}

tavern_events := []string{"intoxicated person knocks over lamp, starts fire","Someone falls out of their seat, but were poisoned","a storm outside blows down a tree falling onto the bar","a storm hits, tree branches break into the window","someone walks in dressed as a horse",
                "someone plays game with the players, cheats the game","a man in tatter clothes is asking people for a meal, the bartender is trying tokick him out","The bar is racist to Elves","a few tables of people are racist to gnomes","a table of non human gang members pick on a human",
                "An investigator is studying his case alone in the corner","someone starts reading from a book to summon a spined devil","a wealthy man tries to buy the bar a round and leave before paying","Bar fight! Love Rivakry","Bar fight! Accidental insult","Bar fight! punching someone on accident",
                "Bar fight! 2 gangs","Bar fight! money"}

tavern_traits := []string{"high class", "shit hole", "gang hangout", "running out of supplies", "average town inn", "secret club running",
                 "well furnished", "big", "small","busy", "ghost town", "secret monster/s hiding ","brawling arena", "monster fighting arena","gambling tables/ machines","haunted",
                 "multidimentional inn(portal inside)","completely run by ","everything shrinks", "huge tavern","full of advanced machinery","everyone inside gets a magic cantrip ability","food eating contest","drinking contests",
                 "pay or fight (pits, portals, teleport etc)","underground rave club","huge tavern /libary. center point of knowlege","suspended over a giant hole in the earth","built around a cave, adventures never come back out","underwater","Band Playing","Glatiator Arena",
                 "jousting arena","Casino","Dog Racing","Horse Racing","Drunk athletic games to win gold","Built In Theater","giant floating ghost ship","head trophies line the walls, adventures come to show off the best trophies","tavern run by a gnome and his clockwork robots/ inventions"}

  num_of_food := rand.Intn(max- 3) + 3

  fmt.Println(tavern_name)
  for i := 0; i < num_of_food; i++{
    fmt.Println("\tFood:\t", tavern_food[rand.Intn(len(tavern_food))])
    fmt.Println("\tDrink:\t", tavern_drinks[rand.Intn(len(tavern_drinks))])
  }
  fmt.Println("Random Event: \t", tavern_events[rand.Intn(len(tavern_events))])
  fmt.Println("Tavern Trait: \t", tavern_traits[rand.Intn(len(tavern_traits))])
}


func gen_taverns(number int,size int) {
  for i := 0; i < number; i++{
    gen_tavern(size*4)
    fmt.Println("\n\n")
  }
}


func gen_city(size int){
  city_name := gen_name() + gen_name()
  fmt.Println(city_name)
  fmt.Println("===========TAVERNS============\n\n")
  for i := 0; i < size; i++{
    gen_tavern(size*4)
    fmt.Println("\n\n")
  }
  fmt.Println("==============STORES===============\n\n")
  for i := 0; i < size*3; i++{
    gen_store()
    fmt.Println("\n\n")
  }

}

func gen_cities(number int, size int){
  for i:= 0; i < number; i++{
    gen_city(size)
  }
}

func main(){
    //Randomize \/\/ Turn off for testing
    rand.Seed(time.Now().UnixNano())
    var input string
    welcome := `

          Welcome to the DND rnadom Generator
                type "name" for a random name/s
                type "npc" for a random NPC/s
                type "store" for a random store/s
                type "tavern" for a random tavern/s
                type "city" for a random city/cities
                type "exit" to quit `

    fmt.Println(welcome)

    for fmt.Scanln(&input); input != "exit"; fmt.Scanln(&input){

      if input == "name"{
        var number int
        fmt.Println("How Many Names?: ")
        fmt.Scanln(&number)
        gen_names(number)

      } else if input == "npc"{
        var number int
        fmt.Println("Number of NPCs?: ")
        fmt.Scanln(&number)
        gen_npcs(number)
      }else if input == "store"{
        var number int
        fmt.Println("Number of Stores?: ")
        fmt.Scanln(&number)
        gen_stores(number)

      }else if input == "tavern"{
        var size int
        var number int
        fmt.Println("What Size? (1,2,3): ")
        fmt.Scanln(&size)
        fmt.Println("how many taverns?: ")
        fmt.Scanln(&number)
        gen_taverns(number,size)

      }else if input == "city"{
        var size int
        var number int
        fmt.Println("What Size? (1,2,3): ")
        fmt.Scanln(&size)
        fmt.Println("how many cities?: ")
        fmt.Scanln(&number)
        gen_cities(number,size)

      }else{
        fmt.Println("Sorry. I did not catch that...")
      }

      fmt.Println(welcome)
    }
}
