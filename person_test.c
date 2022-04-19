/*
 * person_test.c
 * Copyright (C) 2019 Tim Hughes
 *
 * Distributed under terms of the MIT license.
 */

#include <stdio.h>
#include "person.h"
#include "Windows.h"

int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR lpCmdLine, int nShowCmd)
{
    APerson *person = get_person("Tim", "Tim Highes");
    printf("Hello C world: My name is %s, %s.\n", person->name, person->long_name);
    return 0;
}