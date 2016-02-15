/*
Xmms2Go -- A golang binding to libxmmsclient.
Copyright (C) 2016  TonyChyi <tonychee1989@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/*
Package xmms2Go is a Go binding for libxmmsclient.
It's easy to use and friendly for go developers.
Just import me to use.

    package main
    import(
        "github.com/tonychee7000/xmms2go"
        "os"
        "fmt"
    )

    func main(){
        x := xmms2go.NewXmms2Client("test")
        err := x.Connect(os.Getenv("XMMS_PATH"))
        // According to the documents of xmms2, some resources
        // should be released. So Unref() is necessary.
        defer x.Unref()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        err := x.Play()
        if err != nil {
            fmt.Println(err)
        }
    }
*/
package xmms2go
