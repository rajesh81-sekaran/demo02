package sub

import "strconv"
import "fmt"
import "context"
import "crypto/rsa"
import "crypto/rand"
import "github.com/google/go-github/github"

func Foo() {
    client := github.NewClient(nil)

    // list public repositories for org "github"
    opt := &github.RepositoryListByOrgOptions{Type: "public"}
    repos, _, _ := client.Repositories.ListByOrg(context.Background(), "github", opt)

    fmt.Println("Hello from sub! Num repos: " + strconv.Itoa(len(repos)))

    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        panic(err)
    }

    var plaintext = []byte("Lorem ipsum dolor sit amet")
    object, err := encrypter.Encrypt(plaintext)
    if err != nil {
        panic(err)
    }

    fmt.Println(object)
}
