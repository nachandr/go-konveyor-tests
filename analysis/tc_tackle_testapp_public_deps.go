package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/test/api/identity"
)

var MavenPublic = identity.Mvn

var TackleTestappPublicWithDeps = TC{
	Name:        "Tackle Testapp public with deps",
	Application: TackleTestApp,
	Task:        Analyze,
	WithDeps:    true,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=linux",
		},
	},
	Identities: []api.Identity{
		MavenPublic, // Tackle Testapp public Maven registry expects MAVEN_TESTAPP_USER and MAVEN_TESTAPP_TOKEN env variables.

	},
	Analysis: api.Analysis{
		Effort: 3,
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "Hardcoded IP Address",
				Effort:      1,
				RuleSet:     "discovery-rules",
				Rule:        "hardcoded-ip-address",
				Incidents: []api.Incident{
					{
						File:     "/addon/source/tackle-testapp-public/target/classes/persistence.properties",
						Line:     2,
						Message:  "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
						CodeSnip: "jdbc.url=jdbc:oracle:thin:@10.19.2.93:15",
					},
					{
						File:     "/addon/source/tackle-testapp-public/src/main/resources/persistence.properties",
						Line:     2,
						Message:  "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
						CodeSnip: "jdbc.url=jdbc:oracle:thin:@10.19.2.93:15",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "File system - Java IO",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00001",
				Incidents: []api.Incident{
					{
						File:     "/cache/m2/io/konveyor/demo/configuration-utils/1.0.0/io/konveyor/demo/config/ApplicationConfiguration.java",
						Line:     14,
						Message:  "\n An application running inside a container could lose access to a file in local storage",
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Name:     "io.konveyor.demo.configuration-utils",
				Version:  "1.0.0",
				SHA:      "6b39277183eb4c68ee2caa581e7cf1b6d3441b78",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=internal", "konveyor.io/language=java"},
			},
			{
				Name:     "com.h2database.h2",
				Version:  "2.1.214",
				SHA:      "d5c2005c9e3279201e12d4776c948578b16bf8b2",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.postgresql.postgresql",
				Version:  "42.2.23",
				SHA:      "9cb217a3d5b640567ed7c6e8c11f389613c81c4d",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.checkerframework.checker-qual",
				Version:  "3.5.0",
				Indirect: true,
				SHA:      "2f50520c8abea66fbd8d26e481d3aef5c673b510",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "com.oracle.database.jdbc.ojdbc11",
				Version:  "21.1.0.0",
				SHA:      "133b7b0d14b4f4cd4e76661db4727dac937d6856",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "ch.qos.logback.logback-classic",
				Version:  "1.1.7",
				SHA:      "9865cf6994f9ff13fce0bf93f2054ef6c65bb462",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "ch.qos.logback.logback-core",
				Version:  "1.1.7",
				Indirect: true,
				SHA:      "7873092d39ef741575ca91378a6a21c388363ac8",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.hibernate.validator.hibernate-validator",
				Version:  "6.2.0.Final",
				SHA:      "d6b0760dfffbf379cedd02f715ff4c9a2e215921",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "jakarta.validation.jakarta.validation-api",
				Version:  "2.0.2",
				Indirect: true,
				SHA:      "5eacc6522521f7eacb081f95cee1e231648461e7",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "com.fasterxml.classmate",
				Version:  "1.5.1",
				Indirect: true,
				SHA:      "3fe0bed568c62df5e89f4f174c101eab25345b6c",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.hibernate.hibernate-entitymanager",
				Version:  "5.4.32.Final",
				SHA:      "3f60db4097732960ec792c033dbb7c34f1b9e328",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.jboss.logging.jboss-logging",
				Version:  "3.4.1.Final",
				Indirect: true,
				SHA:      "40fd4d696c55793e996d1ff3c475833f836c2498",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.hibernate.hibernate-core",
				Version:  "5.4.32.Final",
				Indirect: true,
				SHA:      "99a5e10bf455337014c190e141ec631e9ff71663",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.javassist.javassist",
				Version:  "3.27.0-GA",
				Indirect: true,
				SHA:      "f63e6aa899e15eca8fdaa402a79af4c417252213",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "antlr.antlr",
				Version:  "2.7.7",
				Indirect: true,
				SHA:      "83cd2cd674a217ade95a4bb83a8a14f351f48bd0",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.jboss.jandex",
				Version:  "2.2.3.Final",
				Indirect: true,
				SHA:      "d3865101f0666b63586683bd811d754517f331ab",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "javax.activation.javax.activation-api",
				Version:  "1.2.0",
				Indirect: true,
				SHA:      "85262acf3ca9816f9537ca47d5adeabaead7cb16",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "javax.xml.bind.jaxb-api",
				Version:  "2.3.1",
				Indirect: true,
				SHA:      "8531ad5ac454cc2deb9d4d32c40c4d7451939b5d",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.glassfish.jaxb.jaxb-runtime",
				Version:  "2.3.1",
				Indirect: true,
				SHA:      "dd6dda9da676a54c5b36ca2806ff95ee017d8738",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.glassfish.jaxb.txw2",
				Version:  "2.3.1",
				Indirect: true,
				SHA:      "a09d2c48d3285f206fafbffe0e50619284e92126",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "com.sun.istack.istack-commons-runtime",
				Version:  "3.0.7",
				Indirect: true,
				SHA:      "c197c86ceec7318b1284bffb49b54226ca774003",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.jvnet.staxex.stax-ex",
				Version:  "1.8",
				Indirect: true,
				SHA:      "8cc35f73da321c29973191f2cf143d29d26a1df7",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "com.sun.xml.fastinfoset.FastInfoset",
				Version:  "1.2.15",
				Indirect: true,
				SHA:      "bb7b7ec0379982b97c62cd17465cb6d9155f68e8",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.dom4j.dom4j",
				Version:  "2.1.3",
				Indirect: true,
				SHA:      "a75914155a9f5808963170ec20653668a2ffd2fd",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.hibernate.common.hibernate-commons-annotations",
				Version:  "5.1.2.Final",
				Indirect: true,
				SHA:      "e59ffdbc6ad09eeb33507b39ffcf287679a498c8",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "javax.persistence.javax.persistence-api",
				Version:  "2.2",
				Indirect: true,
				SHA:      "25665ac8c0b62f50e6488173233239120fc52c96",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "net.bytebuddy.byte-buddy",
				Version:  "1.10.22",
				Indirect: true,
				SHA:      "ef45d7e2cd1c600d279704f492ed5ce2ceb6cdb5",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.jboss.spec.javax.transaction.jboss-transaction-api_1.2_spec",
				Version:  "1.1.1.Final",
				Indirect: true,
				SHA:      "a8485cab9484dda36e9a8c319e76b5cc18797b58",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.apache.tomcat.tomcat-jdbc",
				Version:  "9.0.46",
				SHA:      "385cb6cb1f6b26c881cd5c1c6ade5f180712ffdc",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.apache.tomcat.tomcat-juli",
				Version:  "9.0.46",
				Indirect: true,
				SHA:      "409b519751e104eab51b4347a0d27bf86a4f3bb1",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.boot.spring-boot-starter-actuator",
				Version:  "2.5.0",
				SHA:      "8fc47befa38bdaa2f2b8f421d8532f03005e2851",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.boot.spring-boot-starter",
				Version:  "2.5.0",
				Indirect: true,
				SHA:      "a910887c01efcc7d12f3f89a7604d436f26eeb90",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.boot.spring-boot",
				Version:  "2.5.0",
				Indirect: true,
				SHA:      "b07513e04ad906ea69ef84293a123cdb83828f06",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.boot.spring-boot-autoconfigure",
				Version:  "2.5.0",
				Indirect: true,
				SHA:      "64c7bbc941c70895621ed613f38dc66b73ea9341",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.boot.spring-boot-starter-logging",
				Version:  "2.5.0",
				Indirect: true,
				SHA:      "22401482ba1c5a1dcd3d33e47295779211b913d8",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.apache.logging.log4j.log4j-to-slf4j",
				Version:  "2.14.1",
				Indirect: true,
				SHA:      "ce8a86a3f50a4304749828ce68e7478cafbc8039",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.apache.logging.log4j.log4j-api",
				Version:  "2.14.1",
				Indirect: true,
				SHA:      "cd8858fbbde69f46bce8db1152c18a43328aae78",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.slf4j.jul-to-slf4j",
				Version:  "1.7.30",
				Indirect: true,
				SHA:      "d58bebff8cbf70ff52b59208586095f467656c30",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "jakarta.annotation.jakarta.annotation-api",
				Version:  "1.3.5",
				Indirect: true,
				SHA:      "59eb84ee0d616332ff44aba065f3888cf002cd2d",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.yaml.snakeyaml",
				Version:  "1.28",
				Indirect: true,
				SHA:      "7cae037c3014350c923776548e71c9feb7a69259",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.boot.spring-boot-actuator-autoconfigure",
				Version:  "2.5.0",
				Indirect: true,
				SHA:      "41956882243e86f8260f649ebdd96597a2ff52a9",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.boot.spring-boot-actuator",
				Version:  "2.5.0",
				Indirect: true,
				SHA:      "e0ac75f1a183f8e6a319a8b03bad1c45d40a2761",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "com.fasterxml.jackson.datatype.jackson-datatype-jsr310",
				Version:  "2.12.3",
				Indirect: true,
				SHA:      "f69c636438dcf19c49960c1fe8901320ab85f989",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "io.micrometer.micrometer-core",
				Version:  "1.7.0",
				Indirect: true,
				SHA:      "bc7dc1605f2099dc3c39156b7f62ac889f54fb67",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.hdrhistogram.HdrHistogram",
				Version:  "2.1.12",
				Indirect: true,
				SHA:      "6eb7552156e0d517ae80cc2247be1427c8d90452",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.latencyutils.LatencyUtils",
				Version:  "2.0.3",
				Indirect: true,
				SHA:      "769c0b82cb2421c8256300e907298a9410a2a3d3",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-web",
				Version:  "5.3.7",
				SHA:      "49e6a8f45e77f14ef16f82c0413254ef493b785f",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-webmvc",
				Version:  "5.3.7",
				SHA:      "8437c7a572177a34607abdaef2f6b8088488f5c0",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-expression",
				Version:  "5.3.7",
				Indirect: true,
				SHA:      "13351fce0a604957cd6a41478ebb54a953a0245e",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-jdbc",
				Version:  "5.3.7",
				SHA:      "5caf72035a9b8a3a09ef82322cd2497aedddc487",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.data.spring-data-jpa",
				Version:  "2.5.1",
				SHA:      "881f7ae140f424b3bdb1b0c27a61b93e0bee9fa5",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.data.spring-data-commons",
				Version:  "2.5.1",
				Indirect: true,
				SHA:      "c950ca1a05e928e9fb75420b4ac07713428e9969",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-orm",
				Version:  "5.3.7",
				Indirect: true,
				SHA:      "f1892fe7a6671348d6546facbd40159b7e6f64a2",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-context",
				Version:  "5.3.7",
				Indirect: true,
				SHA:      "330b3957efdcdebe3550b8e2c5d45a4c25496626",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-aop",
				Version:  "5.3.7",
				Indirect: true,
				SHA:      "b86edd2455f8c4399068c999beb9ea2a9e7f2047",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-tx",
				Version:  "5.3.7",
				Indirect: true,
				SHA:      "98be572c2bf3bd08724363b0bba71bcef59c4739",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-beans",
				Version:  "5.3.7",
				Indirect: true,
				SHA:      "8b1eacd7aaa12f7d173a2f0836d28bd0c1b098fe",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-core",
				Version:  "5.3.7",
				Indirect: true,
				SHA:      "4aad1b62bd347a806fe693c9d67b376a3ad8151c",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-jcl",
				Version:  "5.3.7",
				Indirect: true,
				SHA:      "ccd8bde38bad689737295fa220e1c70680676d72",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.aspectj.aspectjrt",
				Version:  "1.9.6",
				Indirect: true,
				SHA:      "1651849d48659e5703adc2599e694bf67b8c3fc4",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.slf4j.slf4j-api",
				Version:  "1.7.26",
				Indirect: true,
				SHA:      "77100a62c2e6f04b53977b9f541044d7d722693d",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-databind",
				Version:  "2.12.3",
				SHA:      "d6153f8fc60c479ab0f9efb35c034526436a4953",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-annotations",
				Version:  "2.12.3",
				Indirect: true,
				SHA:      "7275513412694a1aafd08c0287f48469fa0e6e17",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-core",
				Version:  "2.12.3",
				SHA:      "deb23fe2a7f2b773e18ced2b50d4acc1df8fa366",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.apache.tomcat.tomcat-servlet-api",
				Version:  "9.0.46",
				SHA:      "8e8a27a3456b71b1da2c8adc902ade71bc91fcb4",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Micrometer", Category: api.Ref{Name: "Integration"}},
		{Name: "Spring DI", Category: api.Ref{Name: "Inversion of Control"}},
		{Name: "Spring Data JPA", Category: api.Ref{Name: "Persistence"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "Spring Web", Category: api.Ref{Name: "Web"}},
		{Name: "Spring DI", Category: api.Ref{Name: "Execute"}},
		{Name: "Micrometer", Category: api.Ref{Name: "Execute"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring Web", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring DI", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring Data JPA", Category: api.Ref{Name: "Embedded"}},
		{Name: "Micrometer", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring Web", Category: api.Ref{Name: "View"}},
		{Name: "Spring Data JPA", Category: api.Ref{Name: "Store"}},
	},
}