function verifica() {
    var nome = "Bruno";
    var senha = "b1234";

    var nomeInput = document.getElementById("nomeInput").value;
    var senhaInput = document.getElementById("senhaInput").value;
    var senha = document.getElementById("senha").value; 

    if (nomeInput === nome && senhaInput === senha) {
      alert("Bem-vindo!");
    } else {
      alert("Acesso negado.");
    
    }
  }

