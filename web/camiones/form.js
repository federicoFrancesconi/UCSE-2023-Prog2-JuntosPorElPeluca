const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

const urlConFiltro = "http://go-app:8080/camiones";

document.addEventListener("DOMContentLoaded", function (event) {
  if (!isUserLogged()) {
    window.location =
      document.location.origin + "/login/login.html?reason=login_required";
  }

  //verifico si tiene el parametro id
  const urlParams = new URLSearchParams(window.location.search);
  const patente = urlParams.get("patente");
  const operacion = urlParams.get("tipo");

  if (patente != "" && patente != null && operacion == "EDITAR") {
    document
      .getElementById("form")
      .addEventListener("submit", function (event) {
        actualizarCamion(event);
      });

    document.getElementById("Patente").value = patente;
    document.getElementById("tituloFormulario").innerHTML = "Editar camion";

    //asigno valor a los input si los tiene
    obtenerCamionPorId(patente);
  } else if (patente != "" && patente != null && operacion == "ELIMINAR") {
    eliminarCamion(patente);
    document.getElementById("tituloFormulario").innerHTML = "Eliminar camion";
    document.getElementById("form").style.display = "none";
  } else {
    document
      .getElementById("form")
      .addEventListener("submit", function (event) {
        guardarCamion(event);
      });

    document.getElementById("tituloFormulario").innerHTML = "Crear camion";
  }
});

function guardarCamion(event) {
  event.preventDefault();

  const data = {
    patente: document.getElementById("Patente").value,
    peso_maximo: parseInt(document.getElementById("PesoMaximo").value),
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    id_creador: "",
    costo_por_kilometro: parseInt(document.getElementById("CostoPorKm").value),
  };

  console.log(JSON.stringify(data));

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoCamion,
    errorCamion
  );
}

function exitoCamion(data) {
  alert("Operacion exitosa");
  window.location = window.location.origin + "/camiones/index.html";
}

function errorCamion(status, body) {
  alert(`Error del servidor: ${body.error}`);
  console.log(body.json());
  throw new Error(status.Error);
}

function actualizarCamion(event) {
  event.preventDefault();

  const data = {
    patente: document.getElementById("Patente").value,
    peso_maximo: parseInt(document.getElementById("PesoMaximo").value),
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    id_creador: "",
    costo_por_kilometro: parseInt(document.getElementById("CostoPorKm").value),
  };

  console.log(JSON.stringify(data));

  makeRequest(
    `${urlConFiltro}`,
    Method.PUT,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoCamion,
    errorCamion
  );

  return false;
}

function eliminarCamion(patente) {
  if (confirm("¿Estás seguro de que deseas eliminar este camión?")) {
    makeRequest(
      `${urlConFiltro}/${patente}`,
      Method.DELETE,
      null,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoCamion,
      errorCamion
    );
  } else {
    window.location = document.location.origin + "/camiones/index.html";
  }
}

function obtenerCamionPorId(patente) {
  var url = `http://go-app:8080/camiones/${patente}`;

  makeRequest(
    url,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerCamion,
    errorObtenerCamion
  );
}

function exitoObtenerCamion(data) {
  document.getElementById("Patente").value = data.patente;
  document.getElementById("PesoMaximo").value = data.peso_maximo;
  document.getElementById("CostoPorKm").value = data.costo_por_kilometro;
}

function errorObtenerCamion(status, body) {
  alert(`Error del servidor: ${body.error}`);
  console.log(body.json());
  throw new Error(status.Error);
}
