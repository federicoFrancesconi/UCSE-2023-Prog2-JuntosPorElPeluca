document.addEventListener("DOMContentLoaded", function (event) {
  document.getElementById("form").addEventListener("submit", function (event) {
    obtenerBeneficioEntreFechas(event);
  });
});

function obtenerBeneficioEntreFechas() {
  var fechaDesde = document.getElementById("FechaDesde").value;
  var fechaHasta = document.getElementById("FechaHasta").value;

  var urlConFiltro = `http://localhost:8080/envios/beneficioEntreFechas?fechaDesde=${fechaDesde}Z&fechaHasta=${fechaHasta}Z`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerBeneficioEntreFechas,
    errorEnvio
  );
}

function exitoObtenerBeneficioEntreFechas(data) {
  document.getElementById("beneficio").innerHTML = data;
}
