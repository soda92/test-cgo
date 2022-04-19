/*
 * person.c
 * Copyright (C) 2019 Tim Highes
 * 
 * Distributed under terms of the MIT license.
 */

#include <stdlib.h>
#include "person.h"

APerson *get_person(const char *name, const char* long_name){
    APerson *person = malloc(sizeof(APerson));
    person->name = name;
    person->long_name = long_name;
    return person;
};
