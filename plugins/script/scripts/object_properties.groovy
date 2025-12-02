result = spring.getBean("flexibleSearchService").search("select pk from {$1}").result.get(0)
result.properties.each { println "$it.key -> $it.value" } 
