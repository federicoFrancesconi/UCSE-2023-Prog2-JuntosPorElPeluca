document.addEventListener("DOMContentLoaded", function (event) {
  dibujarGraficoPedidos();
  dibujarGraficoEnvios();
});

function obtenerBeneficioEntreFechas() {
  var fechaDesde = document.getElementById("FechaDesde").value;
  var fechaHasta = document.getElementById("FechaHasta").value;

  var urlConFiltro = `http://localhost:8080/envios/beneficioEntreFechas`;

  //Si fechaDesde esta vacio, no se agrega al filtro
  if (fechaDesde != "") {
    urlConFiltro += `?fechaDesde=${fechaDesde}`;
  }

  //Si fechaHasta esta vacio, no se agrega al filtro
  if (fechaHasta != "") {
    if (fechaDesde != "") {
      urlConFiltro += `&fechaHasta=${fechaHasta}`;
    } else {
      urlConFiltro += `?fechaHasta=${fechaHasta}`;
    }
  }

  //urlConFiltro = `http://localhost:8080/envios/beneficioEntreFechas?fechaDesde=${fechaDesde}&fechaHasta=${fechaHasta}`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerBeneficioEntreFechas,
    errorGraficos
  );
}

function exitoObtenerBeneficioEntreFechas(data) {
  document.getElementById("beneficio").innerHTML = data.beneficio;
}

function errorGraficos(response) {
  alert(response.Error);
  console.log(response.json());
  throw new Error(response.Error);
}

function dibujarGraficoPedidos() {
  var urlConFiltro = `http://localhost:8080/pedidos/cantidadPorEstado`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerGraficoPedidos,
    errorGraficos
  );
}

function exitoObtenerGraficoPedidos(data) {
  var cantidadPedidos = [];
  var estadoPedidos = [];

  for (let i = 0; i < data.length; i++) {
    const element = data[i];
    cantidadPedidos.push(element.Cantidad);
    estadoPedidos.push(element.Estado);
  }

  const datos = {
    labels: estadoPedidos,
    datasets: [
      {
        data: cantidadPedidos, // Cantidad de pedidos por estado
        backgroundColor: [
          "#FF5733",
          "#FFC300",
          "#33FF57",
          "#339CFF",
          "#FFA500",
        ], // Colores para cada sector del gráfico
      },
    ],
  };

  // Configuración del gráfico
  const config = {
    type: "pie",
    data: datos,
  };

  // Dibuja el gráfico en el elemento canvas con id "chart"
  const ctx = document.getElementById("graficoPedidos").getContext("2d");
  new Chart(ctx, config);
}

function dibujarGraficoEnvios() {
  var urlConFiltro = `http://localhost:8080/envios/cantidadPorEstado`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerGraficoEnvios,
    errorGraficos
  );
}

function exitoObtenerGraficoEnvios(data) {
  var cantidadEnvios = [];
  var estadoEnvios = [];

  for (let i = 0; i < data.length; i++) {
    const element = data[i];
    cantidadEnvios.push(element.Cantidad);
    estadoEnvios.push(element.Estado);
  }

  const datos = {
    labels: estadoEnvios,
    datasets: [
      {
        data: cantidadEnvios, // Cantidad de pedidos por estado
        backgroundColor: [
          "#FF5733",
          "#FFC300",
          "#33FF57",
          "#339CFF",
          "#FFA500",
        ], // Colores para cada sector del gráfico
      },
    ],
  };

  // Configuración del gráfico
  const config = {
    type: "pie",
    data: datos,
  };

  // Dibuja el gráfico en el elemento canvas con id "chart"
  const ctx = document.getElementById("graficoEnvios").getContext("2d");
  new Chart(ctx, config);
}
