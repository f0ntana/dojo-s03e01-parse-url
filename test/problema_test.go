package test

import (
	"testing"

	. "github.com/devmatogrosso/dojo-s03e01-parse-url/lib"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProblema(t *testing.T) {
	var problema *Problema

	Convey("Você deve resolver problema", t, func() {

		problema = NewProblema()

		Convey("Mostra Protocolo corretamente", func() {

			So(problema.Resolver("http://www.devmt.com.br").Protocol, ShouldEqual, "http")
			So(problema.Resolver("ftp://www.devmt.com.br").Protocol, ShouldEqual, "ftp")
			So(problema.Resolver("https://www.devmt.com.br").Protocol, ShouldEqual, "https")
			So(problema.Resolver("www.devmt.com.br").Protocol, ShouldEqual, "")
		})

		Convey("Mostra Host corretamente", func() {

			So(problema.Resolver("http://slack.devmt.com.br").Host, ShouldEqual, "slack")
			So(problema.Resolver("ftp://www.devmt.com.br").Host, ShouldEqual, "www")
			So(problema.Resolver("https://xablau.devmt.com.br").Host, ShouldEqual, "xablau")
			So(problema.Resolver("coxinha.devmt.com.br").Host, ShouldEqual, "coxinha")
		})

		Convey("Mostra Dominio corretamente", func() {

			So(problema.Resolver("http://slack.devmt.com.br").Domain, ShouldEqual, "devmt.com.br")
			// So(problema.Resolver("http://google.com").Domain, ShouldEqual, "google.com") TODO
			So(problema.Resolver("ftp://www.devmt.com.br").Domain, ShouldEqual, "devmt.com.br")
			So(problema.Resolver("https://xablau.devmt.com.br").Domain, ShouldEqual, "devmt.com.br")
			So(problema.Resolver("coxinha.devmt.com.br").Domain, ShouldEqual, "devmt.com.br")
		})

	})

}
