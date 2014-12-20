package main

import (
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    // TODO: Use SQL driver.
    //_ "github.com/lib/pq"
    //"database/sql"
)

/*func openDB() (*sql.DB, error) {
    // TODO: ?sslmode=verify-full
    // TODO: Get database info from environment variables
    db, err := sql.Open("postgres", "postgres://chessapp:password@localhost/chessapp")
    return db, err
}*/

// JSON errors based on the Google style guide
// https://google-styleguide.googlecode.com/svn/trunk/jsoncstyleguide.xml#Reserved_Property_Names_in_the_error_object
type Error struct {
    Domain string `json:"domain"`
    Message string `json:"message"`
    Location string `json:"location"`
    LocationType string `json:"locationType"`
    ExtendedHelp string `json:"extendedHelp"`
    SendReport string `json:"sendReport"`
}

type Errors struct {
    Code int64 `json:"code"`
    Message string `json:"message"`
    Errors []Error `json:"errors"`
}

var httpErrors = map[int64]string{
    200: "OK",
    404: "Not Found",
    500: "Internal Server Error",
}

/*
 * Return a json error message for the given code.
 */
func httpError(w http.ResponseWriter, code int64, err []Error) {
    errors := Errors{
        code,
        httpErrors[code],
        err,
    }
    bytes, e := json.Marshal(errors)
    if (e != nil) {
        bytes = []byte(`{"code":500,message:"Internal Server Error",errors:[{domain:"Data",message:"Fatal Error",location:"",locationType:"",extendedHelp:"",sendReport:""}]}`)
    }
    http.Error(w, string(bytes), 500)
    return
}

func asErrorList(err error) []Error {
    return  []Error{
        Error{
            "Data",
            err.Error(),
            "",
            "",
            "",
            "",
        },
    }
}

type Game struct {
    Id int64
    WhitePlayer int64
    BlackPlayer int64
    Result int64
}

func getGames(writer http.ResponseWriter, request *http.Request) {
    games := []Game{
        Game{
            123,
            10,
            11,
            -1,
        },
        Game{
            124,
            10,
            11,
            -1,
        },
    }

    bytedata, e := json.Marshal(games)
    if (e != nil) {
        httpError(writer, 500, asErrorList(e))
        return
    }

    // Send the response.
    fmt.Fprintf(writer, string(bytedata))
}

func getGame(writer http.ResponseWriter, request *http.Request) {
    vars := mux.Vars(request)
    gameId, idErr := strconv.ParseInt(vars["id"], 10, 64)
    if (idErr != nil) {
        httpError(writer, 404, asErrorList(idErr))
        return
    }

    // TODO: Use database.
    /*db, dbErr := openDB()
    if (dbErr != nil) {
        httpError(w, 500, [1]Error{
            Error{
                "Data",
                "Database connection error.",
                e.Error(),
                "",
                "",
                "",
                ""
        }})
        return
    }*/

    // TODO: Get data from database.
    game := Game{
        gameId,
        10,
        11,
        -1,
    }

    bytedata, e := json.Marshal(game)
    if (e != nil) {
        httpError(writer, 500, asErrorList(e))
        return
    }

    // Send the response.
    fmt.Fprintf(writer, string(bytedata))
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/game/{id}/", getGame)
    router.HandleFunc("/game/", getGames)

    http.Handle("/", router)
    http.ListenAndServe(":8080", nil)
}
