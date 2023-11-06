// Ejemplo de uso
const url = "http://w230847.ferozo.com/tp_prog2/api/account/login";

document.addEventListener("DOMContentLoaded", function (eventDOM) {
  document
    .getElementById("btnIngresar")
    .addEventListener("click", function (eventClick) {
      eventClick.preventDefault();

      const data = {
        grant_type: "password",
        username: document.getElementById("usuario").value,
        password: document.getElementById("password").value,
      };

      makeRequest(
        url,
        Method.POST,
        data,
        ContentType.URL_ENCODED,
        CallType.PUBLIC,
        successFn,
        errorFn
      );

      return false;
    });
});

function successFn(response) {
  console.log("Éxito:", response);
  window.location = document.location.origin + "/web/index.html";
}

function errorFn(status, response) {
  console.log("Falla:", response);
}
