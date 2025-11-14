a = spring.getBean("flexibleSearchService").search("select {pk} from {Language}");
a.metaClass.methods.each { method ->
    println "${method.name}( ${method.parameterTypes*.name.join( ', ' )} ) -> ${method.returnType.name} "
}
