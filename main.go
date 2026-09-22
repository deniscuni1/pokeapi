package main
import "fmt"
import "bufio"
import "io"
import "os"
import "strings"
import "net/http"
import "encoding/json"
type config struct {
	commandRegistry map[string]cliCommand
	Next string 
	Previous string
}
type cliCommand struct {
	name string
	description string 
	callback func(*config) error
}
type Results struct {
	Name string
	Url string
}
type Location struct {
	Count int
	Next string
	Previous string
	Results []Results
}
func commandExit(*config) error {
	fmt.Print("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func commandHelp(cfg *config) error {
	fmt.Print("\nUse exit to exit the programme,\nHaven't implemented anything else yet ;p")
	return nil
}
func commandMap(cfg *config) error {
	if cfg.Next != "" {
		resp, err := http.Get(cfg.Next)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		map_of_area := Location{}
		err = json.Unmarshal(data, &map_of_area)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		for i:=0;i<len(map_of_area.Results);i++ {
			fmt.Println(map_of_area.Results[i].Name)
		}
		cfg.Next = map_of_area.Next
		cfg.Previous = map_of_area.Previous
		return nil
	} else {
		resp, err := http.Get("https://pokeapi.co/api/v2/location-area/")
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		map_of_area := Location{}
		err = json.Unmarshal(data, &map_of_area)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		for i:=0;i<len(map_of_area.Results);i++ {
			fmt.Println(map_of_area.Results[i].Name)
			
		}
		cfg.Next = map_of_area.Next
		cfg.Previous = map_of_area.Previous
		return nil
	}
}	
func commandMb(cfg *config) error {
	if cfg.Previous != "" {
		resp, err := http.Get(cfg.Previous)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		map_of_area := Location{}
		err = json.Unmarshal(data, &map_of_area)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		for i:=0;i<len(map_of_area.Results);i++{
			fmt.Println(map_of_area.Results[i].Name)
		}
		cfg.Next = map_of_area.Next
		cfg.Previous = map_of_area.Previous
		return nil
	} else {
		fmt.Println("No previous location")
		return nil
	}
}
var configr = map[string]cliCommand{
	"exit": {
		name: "exit",
		description:"Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description:"Explains how the REPL works",
		callback: commandHelp,
	},
	"map": {
		name: "map",
		description:"Shows 20 location areas in the pokemon world",
		callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description:"Returns to previous location",
		callback: commandMb,
	},
}
var cfg = &config{
				commandRegistry: configr,
				Next: "",
				Previous: "",
			}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to the Pokedex!\n")
	for ;true; {
		fmt.Print("\nPokedex > ")
		if scanner.Scan() == true {
			words := scanner.Text()
			list_of_words := strings.Fields(words)
			if len(list_of_words)==0 {
				continue
			}
			val, exists := cfg.commandRegistry[list_of_words[0]]
			if exists != false {
				err := val.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Print("Unknown command")
			}
		}
	}
}
