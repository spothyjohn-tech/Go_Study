package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// For task number 9
	TypeSingle = "single"
	TypeDouble = "Double"
	TypeSuite  = "Suite"

	StatusFree        = "Free"
	StatusBooked      = "Booked"
	StatusMaintenance = "Maintenance"
	// For task number 12
	bin = 2
	dec = 10
	hex = 16
)

type Order struct {
	ID           int
	Items        []int
	Total        float64
	Adress       string
	isCompleated bool
}

type LogEntry struct {
	IP_Adress string
	HTTP_Code int
	TimePoint time.Time
}

// For task 9
type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

// For task 9
type HotelRoom struct {
	typeRoom   string
	statusRoom string
	priceRoom  float64
}

// For task 10
type TextProperties struct {
	CountSymbols   int
	CountWords     int
	CountSentences int
}

// For task 11
type Product struct {
	Name     string
	Category string
	Price    float64
}

// For task 14
type InventoryItem struct {
	Name        string
	Weight      float64
	IsQuestItem bool
}

// For task 15
type Movie struct {
	Title  string
	Year   int
	Rating float64
	Genres []string
}
// For task 16
type SensorData struct {
	SensorID    string
	Temperature float64
	Humidity    float64
	Timestamp   time.Time
}

func NewOrder(order *map[int]Order) {
	var lastIntKey int
	var Items []int
	var Total float64
	var Adress string
	for k := range *order {
		if k > lastIntKey {
			lastIntKey = k + 1
		}
	}
	for {
		fmt.Print("Введите ID предметов:  (когда закончите введите -1)")
		var Item int
		fmt.Scanln(&Item)
		if Item == -1 {
			break
		} else if Item <= 0 {
			fmt.Print("Такого ID нет")
		} else {
			Items = append(Items, Item)
		}
	}
	for {
		fmt.Print("Введите Цену:")
		fmt.Scanln(&Total)
		if Total > 0 {
			break
		}
	}
	for {
		fmt.Print("Введите адрес:")
		fmt.Scanln(&Adress)
		if Adress != "" {
			break
		}
	}

	orderValue := Order{
		ID:           lastIntKey,
		Items:        Items,
		Total:        Total,
		Adress:       Adress,
		isCompleated: false,
	}

	(*order)[lastIntKey] = orderValue

	fmt.Println("Заказ добавлен")
}

func validateUser(name string, age int, email string) error {
	if name != "" && len(name) > 50 {
		return errors.New("Длина имени не должна быть больше 50 и он не может быть пустым")
	}
	if age < 18 && age > 120 {
		return errors.New("Возраст должен быть между 18 и 120")
	}
	IndexCheck := strings.Index(email, "@")
	if IndexCheck == -1 {

		return errors.New("email должен содержать знак @")
	}
	return nil
}

func CountVotes(Votes []string) {
	var VoteBoris int
	var VoteAnna int
	var VoteVictor int
	var AllVotes float64
	var j string
	for _, j = range Votes {
		if j == "Борис" {
			VoteBoris++
		}
		if j == "Анна" {
			VoteAnna++
		}
		if j == "Виктор" {
			VoteVictor++
		}
	}
	fmt.Printf("Проголосовавших за Анну: %d\n", VoteAnna)
	fmt.Printf("Проголосовавших за Виктора: %d\n", VoteVictor)
	fmt.Printf("Проголосовавших за Бориса: %d\n", VoteBoris)
	AllVotes = float64(VoteBoris+VoteAnna+VoteVictor) / 100
	fmt.Printf("Процент проголосовавших за Анну: %f\n", AllVotes*float64(VoteAnna))
	fmt.Printf("Процент проголосовавших за Виктора: %f\n", AllVotes*float64(VoteVictor))
	fmt.Printf("Процент проголосовавших за Бориса: %f\n", AllVotes*float64(VoteBoris))
}

func CheckSalary(Employees []Employee) (float64, float64) {
	var sumSalary float64
	for _, value := range Employees {
		sumSalary += value.Salary
	}
	AVGsumSalary := sumSalary / float64(len(Employees))
	return sumSalary, AVGsumSalary
}

func SortLogEntries(LogEntries []LogEntry) []LogEntry {
	var SortLogEntries []LogEntry
	for _, value := range LogEntries {
		if value.HTTP_Code >= 400 && value.HTTP_Code <= 599 {
			SortLogEntries = append(SortLogEntries, value)
		}
	}
	return SortLogEntries
}

func BookedRoom(reservRoom *map[string]HotelRoom) {
	var TakeRoom string
	fmt.Println("(Выберите комнату) Свободные комнаты:")
	for i, room := range *reservRoom {
		if room.statusRoom == StatusFree {
			fmt.Println(i)
		}
	}
	for {
		fmt.Scanln(&TakeRoom)
		if TakeRoom != "" {
			break
		}
	}
	if room, exists := (*reservRoom)[TakeRoom]; exists {
		if room.statusRoom == StatusFree {
			room.statusRoom = StatusBooked
			(*reservRoom)[TakeRoom] = room
			fmt.Println("Комната: " + TakeRoom + " была забронирована")
		}
	} else {
		fmt.Println("Такой комнаты для бронирования нет")
	}

}

func textStats(text string) TextProperties {
	var TextStatistic TextProperties
	TextStatistic.CountSymbols = len([]rune(text))
	TextStatistic.CountWords = len(strings.Fields(text))
	for _, Symbol := range text {
		if Symbol == '!' || Symbol == '.' || Symbol == '?' {
			TextStatistic.CountSentences += 1
		}
	}
	return TextStatistic
}

func filterProducts(products []Product, maxPrice float64, category string) []Product {
	var FilterProducts []Product
	for _, product := range products {
		if product.Category == category {
			if product.Price < maxPrice {
				FilterProducts = append(FilterProducts, product)
			}
		}
	}
	return FilterProducts
}

func convertNumber(number string, OriginalNumbSystem int, MoveInSystem int) string {
	result, err := strconv.ParseInt(number, OriginalNumbSystem, 64)
	if err != nil {
		return "wrong number"
	} else {
		str := strconv.FormatInt(result, MoveInSystem)
		fmt.Printf("Перевод числа "+number+" из %d - ой системы в %d - ую систему будет равно: "+str, OriginalNumbSystem, MoveInSystem)
		return str
	}
}

func AddExpenses(TypeExpense string, SumExpanse float64, Expenses *map[string]float64) {
	(*Expenses)[TypeExpense] += SumExpanse
}

func SumAllWeight(PlayerInventory []InventoryItem) float64 {
	var AllWeight float64
	for _, item := range PlayerInventory {
		AllWeight += item.Weight
	}
	return AllWeight
}

func FoundFilmWithBigRating(Movies []Movie) {
	var FilmName string
	var Reting float64
	for _, film := range Movies {
		if FilmName == "" {
			FilmName = film.Title
			Reting = film.Rating
		} else {
			if Reting < film.Rating {
				FilmName = film.Title
				Reting = film.Rating
			}
		}
	}
	fmt.Printf("Фильм с самым высоким рейтингом: "+FilmName+" рейтинг составляет: %f \n", Reting)
	fmt.Println("P.S. Никому это не изменить!!!!")
}

func AddGenre(Movies []Movie) {
	var FilmName string
	var GenreForAdd string
	var foundFilm bool
	fmt.Println("Выберите фильм чтобы добавить ему жанр")
	fmt.Println("Список фильмов:")
	for _, film := range Movies {
		fmt.Println(film)
	}
	fmt.Println("")
	fmt.Println("Введите к какому фильму хотите добавить жанр:")
	fmt.Scanln(&FilmName)
	fmt.Println("Введите какой жанр хотите добавить:")
	fmt.Scanln(&GenreForAdd)
	for i, film := range Movies {
		if film.Title == FilmName {
			Movies[i].Genres = append(Movies[i].Genres, GenreForAdd)
			foundFilm = true
		}
	}
	if !foundFilm {
		fmt.Println("Данный фильм не был найден")
	} else {
		for _, film := range Movies {
			if film.Title == FilmName {
				fmt.Println("Фильм:")
				fmt.Println(film.Title)
				fmt.Println("Жанры:")
				fmt.Println(film.Genres)
			}
		}
	}
}

func FoundFilmGenre(Movies []Movie) {
	fmt.Println("---------------Поиск фильмов по жанрам---------------")
	fmt.Println("--------------------Список жанров:--------------------")
	var ListGenres []string
	var ChoiceGenre string
	var isRepeat bool
	for _, film := range Movies {
		for _, Genre := range film.Genres {
			isRepeat = false
			for _, GenreInList := range ListGenres {
				if GenreInList == Genre {
					isRepeat = true
				}
			}
			if !isRepeat {
				ListGenres = append(ListGenres, Genre)
			}
		}
	}
	for _, GenreInList := range ListGenres {
		fmt.Println(GenreInList)
	}
	fmt.Println("")
	fmt.Println("Введите жанр:")
	fmt.Scanln(&ChoiceGenre)
	fmt.Println("--------------------Список фильмов с данным жанром:--------------------")
	for _, film := range Movies {
		for _, Genre := range film.Genres {
			if Genre == ChoiceGenre {
				fmt.Println(film.Title)
				break
			}
		}
	}
}

func AverageTemperature(information []SensorData) float64 {
	var sum float64
	var count int
	for _, sensor := range information {
		sum += sensor.Temperature
		count++
	}
	return sum / float64(count)
}

func main() {
	var Exit bool = false
	var TaskManage int
	for {
		fmt.Println("")
		if Exit {
			break
		}
		fmt.Println("Выберите задание (введите значение от 1-16) (-1 чтобы выйти): ")
		fmt.Scanln(&TaskManage)
		switch TaskManage {
		case 1:
			PriceDays := map[string]int{
				"ПН": 2100,
				"ВТ": 2100,
				"СР": 2100,
				"ЧТ": 2100,
				"ПТ": 2850,
				"СБ": 2850,
				"ВС": 2850,
			}
			price := PriceDays["ВТ"] + PriceDays["СР"] + PriceDays["ЧТ"] + PriceDays["ПТ"] + PriceDays["СБ"] + PriceDays["ВС"] + PriceDays["ЧТ"] + PriceDays["ПТ"]
			fmt.Println(price)
		case 2:
			var WeightBagageMain float64
			var WeightBagageHand float64
			var WeightBagageHandDop float64
			for {
				fmt.Print("Введите вес основного багажа: ")
				fmt.Scanln(&WeightBagageMain)
				fmt.Print("Введите вес ручной клади: ")
				fmt.Scanln(&WeightBagageHand)
				fmt.Print("Введите вес доп. ручной клади: ")
				fmt.Scanln(&WeightBagageHandDop)
				if WeightBagageMain >= 0 && WeightBagageHand >= 0 && WeightBagageHandDop >= 0 {
					break
				}
			}
			fmt.Printf("Общий вес багажа: %f", WeightBagageMain+WeightBagageHand+WeightBagageHandDop)
			fmt.Println()
		case 3:
			Orders := map[int]Order{
				1: {
					ID:           1,
					Items:        []int{2, 3, 4},
					Total:        200.5,
					Adress:       "ул. Ленина 25",
					isCompleated: false,
				},
			}
			NewOrder(&Orders)
			fmt.Println(Orders)
		case 4:
			var CondidatsVotes = []string{"Анна", "Борис", "Виктор", "Борис", "Виктор", "Борис", "Виктор", "Анна", "Анна", "Анна"}
			CountVotes(CondidatsVotes)
		case 5:
			err := validateUser("Грыгорий", 20, "email@email.com")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Ошибок в регистрации не было обнаружено")
			}
		case 6:
			var Tags = [][]string{
				{"go", "backend"},
				{"git", "go", "tools"},
			}
			var isRepeat bool
			UniqueTags := make(map[string]bool)
			for _, ListTags := range Tags {
				for _, tag := range ListTags {
					isRepeat = false
					for UniqTag := range UniqueTags {
						if UniqTag == tag {
							isRepeat = true
						}
					}
					if !isRepeat {
						UniqueTags[tag] = true
					}
				}
			}
			fmt.Println("Все уникальные тэги:")
			for UniqTag := range UniqueTags {
				fmt.Println(UniqTag)
			}
		case 7:
			Employees := []Employee{
				{
					ID:       1,
					Name:     "Grigoriy",
					Position: "Слесарь",
					Salary:   35000,
				},
				{
					ID:       2,
					Name:     "Alex",
					Position: "Программист",
					Salary:   25000,
				},
				{
					ID:       3,
					Name:     "Alena",
					Position: "Бухгалтер",
					Salary:   15000,
				},
			}
			SumSalary, AvgSumSalary := CheckSalary(Employees)
			fmt.Printf("Общая получаемая сумма рабочими: %f Средняя получаемая сумма: %f \n", SumSalary, AvgSumSalary)
		case 8:
			LogEntries := []LogEntry{
				{
					IP_Adress: "198.5.3.1",
					HTTP_Code: 405,
					TimePoint: time.Now(),
				},
				{
					IP_Adress: "195.5.1.1",
					HTTP_Code: 129,
					TimePoint: time.Now(),
				},
				{
					IP_Adress: "178.5.3.1",
					HTTP_Code: 506,
					TimePoint: time.Now(),
				},
			}
			fmt.Println("Логи до этого:")
			for _, value := range LogEntries {
				fmt.Println(value)
			}
			var SortLogEntries []LogEntry = SortLogEntries(LogEntries)
			fmt.Println("Отсортированные логи:")
			for _, value := range SortLogEntries {
				fmt.Println(value)
			}
		case 9:
			reservRoom := map[string]HotelRoom{
				"109": {
					typeRoom:   TypeSingle,
					statusRoom: StatusFree,
					priceRoom:  305,
				},
				"111": {
					typeRoom:   TypeDouble,
					statusRoom: StatusBooked,
					priceRoom:  400,
				},
				"101": {
					typeRoom:   TypeSuite,
					statusRoom: StatusMaintenance,
					priceRoom:  500,
				},
			}
			BookedRoom(&reservRoom)
			fmt.Println("Комнаты и их статус:")
			for i, room := range reservRoom {
				fmt.Println(i + ":" + " " + room.typeRoom + " " + room.statusRoom)

			}
		case 10:
			TextStatistic := textStats("Предложение. Предложение!")
			fmt.Println(TextStatistic)
		case 11:
			Products := []Product{
				{
					Name:     "Хлеб",
					Category: "Хлебобулочные",
					Price:    30,
				},
				{
					Name:     "Булочка",
					Category: "Хлебобулочные",
					Price:    20,
				},
				{
					Name:     "Батон",
					Category: "Хлебобулочные",
					Price:    40,
				},
				{
					Name:     "Шоколадка",
					Category: "Сладкое",
					Price:    100,
				},
				{
					Name:     "Конфеты",
					Category: "Сладкое",
					Price:    150,
				},
				{
					Name:     "Яблоко",
					Category: "Фрукты",
					Price:    40,
				},
				{
					Name:     "Грушка",
					Category: "Фрукты",
					Price:    50,
				},
				{
					Name:     "Бананы",
					Category: "Фрукты",
					Price:    200,
				},
			}
			FilterProducts := filterProducts(Products, 100, "Фрукты")
			fmt.Println("Отфильтровованные продукты:")
			for _, FiltProduct := range FilterProducts {
				fmt.Println(FiltProduct)
			}
		case 12:
			var InputNumber string
			var SystemNumber int
			var SystemNumberConv int
			fmt.Println("Перевод из одной системы в другую")
			fmt.Println("______________________________________________")
			fmt.Println("Введите число:")
			fmt.Scanln(&InputNumber)
			fmt.Println("В какой он системе счисления? (2, 10, 16): ")
			fmt.Scanln(&SystemNumber)
			fmt.Println("В какую его перевести? (2, 10, 16): ")
			fmt.Scanln(&SystemNumberConv)
			convertNumber(InputNumber, SystemNumber, SystemNumberConv)
		case 13:
			var SumExpense float64
			var TypeExpense string
			expenses := map[string]float64{
				"Food":    15000,
				"Vehicle": 5000,
				"Fun":     3000,
			}
			for {
				fmt.Println("Введите категорию траты (Food, Vehicle, Fun): ")
				fmt.Scanln(&TypeExpense)
				fmt.Println("Введите сколько потратили: ")
				fmt.Scanln(&SumExpense)
				if TypeExpense == "Food" || TypeExpense == "Vehicle" || TypeExpense == "Fun" && SumExpense >= 0 {
					break
				}
			}
			AddExpenses(TypeExpense, SumExpense, &expenses)
			fmt.Println("Сумма всех трат:")
			fmt.Println(expenses["Food"] + expenses["Vehicle"] + expenses["Fun"])
		case 14:
			inventory := []InventoryItem{
				{
					Name:        "Sword",
					Weight:      4.5,
					IsQuestItem: false,
				},
				{
					Name:        "Axe",
					Weight:      5,
					IsQuestItem: false,
				},
				{
					Name:        "Magic_Ball",
					Weight:      1,
					IsQuestItem: true,
				},
				{
					Name:        "Chair",
					Weight:      2,
					IsQuestItem: true,
				},
				{
					Name:        "Very_Need_Item",
					Weight:      100,
					IsQuestItem: false,
				},
			}
			SumWeight := SumAllWeight(inventory)
			fmt.Printf("Сумма веса всех ваших вещей в инвенторе равен: %f (похоже у вас есть очень важный предмет)", SumWeight)
		case 15:
			var Choice int
			Movies := []Movie{
				{
					Title:  "Sword",
					Year:   1990,
					Rating: 7.5,
					Genres: []string{"War", "SuperWar"},
				},
				{
					Title:  "Старый_мужик_есть_суп",
					Year:   1,
					Rating: 10,
					Genres: []string{"Боевик", "Fun"},
				},
				{
					Title:  "Яндере_нашла_кошку_которая_оказалась_Джон_Уикот",
					Year:   2025,
					Rating: 1,
					Genres: []string{"isVeryBad", "Anime"},
				},
				{
					Title:  "John Wick",
					Year:   2014,
					Rating: 1000,
					Genres: []string{"Лучший", "ЛУЧШИЙлучшийЛУЧШИЙ"},
				},
				{
					Title:  "Joker",
					Year:   2019,
					Rating: 5,
					Genres: []string{"5_клоунов_из_десяти"},
				},
			}
			fmt.Println("Выберите что хотите сделать:")
			fmt.Println(" 1. Посмотреть фильм с самым высоким рейтингом \n 2. Добавить жанр фильму \n 3. Посмотреть какие есть фильмы по жанру")
			fmt.Scanln(&Choice)
			if Choice == 1 {
				FoundFilmWithBigRating(Movies)
			} else if Choice == 2 {
				AddGenre(Movies)
			} else if Choice == 3 {
				FoundFilmGenre(Movies)
			} else {
				fmt.Println("Такого варианта нет")
			}
		case 16:
			readings := []SensorData{
				{
					SensorID:    "Sensor1",
					Temperature: 22.5,
					Humidity:    45.0,
					Timestamp:   time.Date(2025, 6, 9, 8, 0, 0, 0, time.UTC),
				},
				{
					SensorID:    "Sensor2",
					Temperature: 23.1,
					Humidity:    50.0,
					Timestamp:   time.Date(2025, 6, 9, 9, 0, 0, 0, time.UTC),
				},
				{
					SensorID:    "Sensor3",
					Temperature: 21.8,
					Humidity:    47.0,
					Timestamp:   time.Date(2025, 6, 9, 10, 0, 0, 0, time.UTC),
				},
				{
					SensorID:    "Sensor4",
					Temperature: 22.0,
					Humidity:    46.0,
					Timestamp:   time.Date(2025, 6, 9, 11, 0, 0, 0, time.UTC),
				},
				{
					SensorID:    "Sensor5",
					Temperature: 23.5,
					Humidity:    52.0,
					Timestamp:   time.Date(2025, 6, 9, 12, 0, 0, 0, time.UTC),
				},
			}
			AVGTemperature := AverageTemperature(readings)
			fmt.Printf("Средняя температкра дома равна: %f", AVGTemperature)
		case -1:
			Exit = true
		default:
			fmt.Println("Данного значения нет")
		}
	}
}
