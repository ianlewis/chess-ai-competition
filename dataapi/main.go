package main

import (
    "fmt"
    "net/http"
    "encoding/json"
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
    domain string
    message string
    location string
    locationType string
    extendedHelp string
    sendReport string
}

type Errors struct {
    code int64
    message string
    errors []Error
}

var httpErrors = map[int64]string{
    200: "OK",
    404: "Not Found",
    500: "Internal Server Error",
}

func httpError(w http.ResponseWriter, code int64, err []Error) {
    // Return a json error message.
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

type Game struct {
    Id int64
    WhitePlayer int64
    BlackPlayer int64
    IsPlaying bool
    Result int64
}

func getGame(w http.ResponseWriter, r *http.Request) {
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
        123,
        10,
        11,
        true,
        -1,
    }

    bytedata, e := json.Marshal(game)
    if (e != nil) {
        httpError(w, 500, []Error{
            Error{
                "Data",
                "Fatal json.Marshall() error.",
                e.Error(),
                "",
                "",
                "",
            },
        })
        return
    }

    // Send the response.
    fmt.Fprintf(w, string(bytedata))
}

func main() {
    http.HandleFunc("/", getGame)
    http.ListenAndServe(":8080", nil)
}
