document.addEventListener("DOMContentLoaded", function (event) {
  if (!isUserLogged()) {
    window.location =
      document.location.origin + "/web/login/login.html?reason=login_required";
  }

  //verifico si tiene el parametro id
  const urlParams = new URLSearchParams(window.location.search);
  const idEnvio = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (idEnvio != "" && idEnvio != null && operacion == "INICIAR") {
    iniciarViaje(idEnvio);
    document.getElementById("form").style.display = "none";
    document.getElementById("listaPedidos").style.display = "none";
  } else if (idEnvio != "" && idEnvio != null && operacion == "FINALIZAR") {
    finalizarViaje(idEnvio);
    document.getElementById("form").style.display = "none";
    document.getElementById("listaPedidos").style.display = "none";
  } else {
    document
      .getElementById("buttonSave")
      .addEventListener("click", function (event) {
        guardarEnvio(event);
      });

    obtenerPedidos();
  }
});

//obtiene los pedidos para mostrar en el form de crear
function obtenerPedidos() {
  const urlConFiltro = `http://localhost:8080/pedidos?estado=Aceptado`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerPedidosEnvio,
    errorEnvio
  );
}

var productosPedidos = [];

function exitoObtenerPedidosEnvio(data) {
  const elementosTable = document //tabla en la que se colocan los envios que se obtienen
    .getElementById("tablePedidos")
    .querySelector("tbody");

  console.log(elementosTable);

  data.forEach((elemento) => {
    const row = document.createElement("tr"); //crear una fila

    row.innerHTML = ` 
                  <td><input type="checkbox" class="pedido-checkbox"></td>
                  <td>${elemento.id}</td>
                  <td>${elemento.ciudad_destino}</td>
                  <td>${elemento.estado}</td>
                  <td>${elemento.fecha_creacion}</td>
                  <td>${elemento.fecha_ultima_actualizacion}</td>
                  <td>${elemento.id_creador}</td>
                 `;
    const nuevoObjetoPedido = {
      id: elemento.id,
      productos: elemento.productos_elegidos,
    };
    productosPedidos.push(nuevoObjetoPedido);
    elementosTable.appendChild(row);
  });
}

function obtenerPedidosArray() {
  var PedidosSeleccionados = [];
  let checkboxes = document.querySelectorAll(".pedido-checkbox");
  checkboxes.forEach(function (checkbox) {
    if (checkbox.checked) {
      // Agregar el producto seleccionado al objeto ProductosSeleccionados
      var tr = checkbox.closest("tr");

      var idPedido = tr.cells[1].textContent;

      PedidosSeleccionados.push(idPedido);
    }
  });

  return PedidosSeleccionados;
}

const urlConFiltro = `http://localhost:8080/envios`;

function guardarEnvio() {
  const pedidosArray = obtenerPedidosArray();

  //armo la data a enviar
  const data = {
    id: "",
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    patente_camion: document.getElementById("PatenteCamion").value,
    paradas: [],
    pedidos: pedidosArray,
    estado: "ADespachar",
  };

  //convierte a json la data
  debugger;
  const json = JSON.stringify(data);
  console.log(json);

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoEnvio,
    errorEnvio
  );
}

function exitoEnvio(data) {
  alert("Operacion exitosa");
  window.location = window.location.origin + "/web/envios/index.html";
}

function errorEnvio(response) {
  alert(`Error del servidor: ${response.error}`);
  console.log(response.json());
  throw new Error(response.Error);
}

function iniciarViaje(id) {
  dataEnvio = {
    id: id,
    estado: "En Ruta",
  };

  if (confirm("¿Estás seguro de que deseas iniciar el viaje?")) {
    makeRequest(
      `${urlConFiltro}/cambiarEstado`,
      Method.PUT,
      dataEnvio,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoEnvio,
      errorEnvio
    );
  } else {
    window.location = document.location.origin + "/web/envios/index.html";
  }
}

function finalizarViaje(id) {
  window.location =
    document.location.origin +
    `/web/envios/nuevaParada.html?id=${id}&tipo=FINALIZAR`;
}
