// Ejemplo de uso
const url = "http://w220066.ferozo.com/tp_prog2/api/account/register";

document.addEventListener("DOMContentLoaded", function (eventDOM) {
  document
    .getElementById("btnRegistrarse")
    .addEventListener("click", function (eventClick) {
      eventClick.preventDefault();

      const inputPassword = document.getElementById("password").value;
      const inputConfirmPassword =
        document.getElementById("ConfirmPassword").value;

      if (inputPassword !== inputConfirmPassword) {
        alert("Password and Confirm Password must be the same");
        return false;
      }

      const data = {
        grant_type: "password",
        email: document.getElementById("email").value,
        password: inputPassword,
        ConfirmPassword: inputConfirmPassword,
        Role: document.getElementById("Role").value,
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
  window.location = "index.html";
}

function errorFn(status, response) {
  console.log("Falla:", response);
}
