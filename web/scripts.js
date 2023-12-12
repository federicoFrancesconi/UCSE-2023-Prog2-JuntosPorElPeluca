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
  var montoFechasMes = [];
  var montoFechasAnio = [];
  var meses = [];
  var anios = [];

  if (data.length === 0) {
    document.getElementById("mensajeSinBeneficio").innerHTML = "No hay beneficios cargados en esas fechas";
    return; // Agregamos un return para salir de la función si no hay datos
  }

  // Procesar datos por mes
  data.meses.forEach((element) => {
    montoFechasMes.push(element.Monto);
    meses.push(element.Nombre);
  });

  // Procesar datos por año
  data.años.forEach((element) => {
    montoFechasAnio.push(element.Monto);
    anios.push(element.Nombre);
  });

  const configuracionBarras = {
    responsive: true,
    scales: {
      y: {
        beginAtZero: true,
      },
    },
  };

  const datosMeses = {
    labels: meses,
    datasets: [
      {
        data: montoFechasMes,
        backgroundColor: ["#FF5733", "#FFC300", "#33FF57", "#339CFF", "#FFA500"],
      },
    ],
  };

  // Obtener el contexto del lienzo de barras de meses
  var contextoBarrasMeses = document.getElementById('graficoBeneficioMes').getContext('2d');

  // Crear el gráfico de barras de meses
  var miGraficoBarrasMeses = new Chart(contextoBarrasMeses, {
    type: 'bar',
    data: datosMeses,
    options: configuracionBarras,
  });

  const datosAnios = {
    labels: anios,
    datasets: [
      {
        data: montoFechasAnio,
        backgroundColor: ["#FF5733", "#FFC300", "#33FF57", "#339CFF", "#FFA500"],
      },
    ],
  };

  // Obtener el contexto del lienzo de barras de años
  var contextoBarrasAnio = document.getElementById('graficoBeneficioAnio').getContext('2d');

  // Crear el gráfico de barras de años
  var miGraficoBarrasAnios = new Chart(contextoBarrasAnio, {
    type: 'bar',
    data: datosAnios,
    options: configuracionBarras,
  });

  new Chart(contextoBarrasMeses, miGraficoBarrasMeses);

  new Chart(contextoBarrasAnio, miGraficoBarrasAnios);

}


function errorGraficos(status, body) {
  alert(body.error);
  console.log(body.json());
  throw new Error(status.Error);
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

  if (data.length == 0) {
    document.getElementById("mensajeSinPedidos").innerHTML = "No hay pedidos cargados";
  }

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

  // Dibuja el gráfico en el elemento canvas 
  const ctx = document.getElementById("graficoPedidosTorta").getContext("2d");
  new Chart(ctx, config);

  // Configuración del gráfico de barras
  var configuracionBarras = {
    responsive: true,
    scales: {
        y: {
            beginAtZero: true
        }
    }
  };

  // Obtener el contexto del lienzo de barras
  var contextoBarras = document.getElementById('graficoPedidosBarra').getContext('2d');

  // Crear el gráfico de barras
  var miGraficoBarras = new Chart(contextoBarras, {
      type: 'bar',
      data: datos,
      options: configuracionBarras
  });

  // Dibuja el gráfico de barras en el elemento canvas 
  new Chart(contextoBarras, miGraficoBarras);
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

  if (data.length == 0) {
    document.getElementById("mensajeSinEnvios").innerHTML = "No hay envios cargados";
    return;
  }

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

  // Configuración del gráfico de torta
  const config = {
    type: "pie",
    data: datos,
  };

  // Dibuja el gráfico de torta en el elemento canvas 
  const ctx = document.getElementById("graficoEnviosTorta").getContext("2d");
  new Chart(ctx, config);

  // Configuración del gráfico de barras
  var configuracionBarras = {
    responsive: true,
    scales: {
        y: {
            beginAtZero: true
        }
    }
  };

  const configBarras = {
    type: 'bar',
      data: datos,
      options: configuracionBarras
  };

  // Dibuja el gráfico de barras en el elemento canvas 
  const ctxBarras = document.getElementById("graficoEnviosBarra").getContext("2d");
  new Chart(ctxBarras, configBarras);
}
