package main

import ("fmt"
        "bufio"
        "os"
        "strconv"
        . "github.com/logrusorgru/aurora"
      )

func main(){
  fmt.Println(Bold(Cyan("\n{#PLUGİNLER.COM#}")), Bold(Green("\n{--MİTHAT GÜNER--}\n")))
  fmt.Println(Bold(Gray("-- | Hesap Makinesi Uygulaması| --\n")))
  fmt.Println(Bold(Brown("\n# Toplama Yap (+)\n# Çıkarma Yap (-)\n# Çarpma Yap (*)\n# Bölme Yap (/)")))
  fmt.Println(Bold(Red("\n\n-- | İşlem Yapmak İstediğiniz Sembolü Girin | --\n")))
  tara := bufio.NewScanner(os.Stdin)
  tara.Scan()
  cikti := tara.Text()

  if(cikti == "+"){
    fmt.Println("\n-- | Toplama İşlemi Yapacağınız İlk Sayıyı Girin : \n")
    tara.Scan()
    ilk_sayi := tara.Text()
    fmt.Println("\n-- | Toplama İşlemi Yapacağınız İkinci Sayıyı Girin : \n")
    tara.Scan()
    iki_sayi := tara.Text()


    int_sayi1, _ := strconv.ParseInt(ilk_sayi, 10, 0)
    int_sayi2, _ := strconv.ParseInt(iki_sayi, 10, 0)

    fmt.Println("\n-- | ",int_sayi1, " + ", int_sayi2, " = " ,int_sayi1 + int_sayi2)
  }else if(cikti == "-"){
    fmt.Println("\n-- | Çıkarma İşlemi Yapacağınız İlk Sayıyı Girin : \n")
    tara.Scan()
    ilk_sayi := tara.Text()
    fmt.Println("\n-- | Çıkarma İşlemi Yapacağınız İkinci Sayıyı Girin : \n")
    tara.Scan()
    iki_sayi := tara.Text()


    int_sayi1, _ := strconv.ParseInt(ilk_sayi, 10, 0)
    int_sayi2, _ := strconv.ParseInt(iki_sayi, 10, 0)

    fmt.Println("\n-- | ",int_sayi1, " - ", int_sayi2, " = " ,int_sayi1 - int_sayi2)
  }else if(cikti == "*"){
    fmt.Println("\n-- | Çarpma İşlemi Yapacağınız İlk Sayıyı Girin : \n")
    tara.Scan()
    ilk_sayi := tara.Text()
    fmt.Println("\n-- | Çarpma İşlemi Yapacağınız İkinci Sayıyı Girin : \n")
    tara.Scan()
    iki_sayi := tara.Text()


    int_sayi1, _ := strconv.ParseInt(ilk_sayi, 10, 0)
    int_sayi2, _ := strconv.ParseInt(iki_sayi, 10, 0)

    fmt.Println("\n-- | ",int_sayi1, " * ", int_sayi2, " = " ,int_sayi1 * int_sayi2)
  }else if(cikti == "/"){
    fmt.Println("\n-- | Bölme İşlemi Yapacağınız İlk Sayıyı Girin : \n")
    tara.Scan()
    ilk_sayi := tara.Text()
    fmt.Println("\n-- | Bölme İşlemi Yapacağınız İkinci Sayıyı Girin : \n")
    tara.Scan()
    iki_sayi := tara.Text()


    int_sayi1, _ := strconv.ParseInt(ilk_sayi, 10, 0)
    int_sayi2, _ := strconv.ParseInt(iki_sayi, 10, 0)

    fmt.Println("\n-- | ",int_sayi1, " / ", int_sayi2, " = " ,int_sayi1 / int_sayi2)
  }
}
