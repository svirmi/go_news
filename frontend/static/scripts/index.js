"use strict";

var go = document.getElementById("button-go");
var results = document.getElementById("results");
var select = document.getElementById("source");

window.onload = () => {
  go.addEventListener("click", () => {
    fetch(`/search/${select.value}`, { mode: "no-cors" })
      .then((response) => {
        debugger;
        if (response.ok) {
          return response.text();
        }
      })
      .then((data) => {
        results.innerHTML = data;
      })
      .catch((err) => {
        console.error(err.message);
      });
  });
};
