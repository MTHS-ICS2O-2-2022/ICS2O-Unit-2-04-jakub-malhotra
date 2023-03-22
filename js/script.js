// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Jakub Malhotra
// Created on: March 2023
// This file contains the JS functions for index.html

"use strict"
/**
 * This function calculates the area of a triangle
 */
function enterClicked() {
  //input
  const base = parseInt(document.getElementById("base-of-triangle").value)
  const height = parseInt(document.getElementById("height-of-triangle").value)

  //process
  const area = (base * height) / 2

  //output
  document.getElementById("area-of-triangle").innerHTML =
    "The area of the triangle is " + area + " cm²."
}