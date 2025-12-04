document.addEventListener("DOMContentLoaded", () => {
  const textInput = document.getElementById("text-input")
  const generateButton = document.getElementById("generate-button")
  const asciiOutput = document.getElementById("ascii-output-text")
  const fontSelect = document.getElementById("font-select")
  const colorOptions = document.querySelectorAll(".color-option")
  const downloadAsciiBtn = document.getElementById("download-ascii")
  const copyAsciiBtn = document.getElementById("copy-ascii")

  let currentColor = "#ffffff" // Couleur par défaut (Blanc)

  // --- Gestion de la sélection de couleur ---
  colorOptions.forEach((option) => {
    option.addEventListener("click", () => {
      // Désélectionner toutes les options
      colorOptions.forEach((opt) => opt.classList.remove("selected"))

      // Sélectionner l'option cliquée
      option.classList.add("selected")

      // Mettre à jour la couleur actuelle
      currentColor = option.getAttribute("data-color")

      // Appliquer la couleur au texte ASCII
      asciiOutput.style.color = currentColor
    })
  })

  // Initialiser la couleur (pour le blanc par défaut)
  const defaultColorOption = document.querySelector('.color-option[data-color="#ffffff"]')
  if (defaultColorOption) {
    defaultColorOption.classList.add("selected")
  }
  asciiOutput.style.color = currentColor

  // Variable pour stocker la police sélectionnée
  let selectedBanner = "standard"

  // --- Gestion de la sélection de police ---
  fontSelect.addEventListener("change", (e) => {
    selectedBanner = e.target.value
  })

  // --- Gestion du bouton GÉNÉRER ---
  generateButton.addEventListener("click", async () => {
    const textToConvert = textInput.value

    if (textToConvert.trim() === "") {
      asciiOutput.textContent = "Veuillez entrer du texte à convertir."
      return
    }

    // Afficher un message de chargement
    asciiOutput.textContent = "Génération en cours..."

    try {
      // Envoyer une requête POST au serveur
      const formData = new FormData()
      formData.append("text", textToConvert)
      formData.append("banner", selectedBanner)

      const response = await fetch("/ascii-art", {
        method: "POST",
        body: formData
      })

      if (!response.ok) {
        throw new Error(`Erreur HTTP: ${response.status}`)
      }

      const asciiArt = await response.text()
      
      // Afficher l'ASCII Art généré
      asciiOutput.textContent = asciiArt
      asciiOutput.style.color = currentColor
    } catch (error) {
      asciiOutput.textContent = `Erreur lors de la génération: ${error.message}`
      console.error("Erreur:", error)
    }
  })

  // --- Gestion des boutons de Téléchargement et Copie ---
  downloadAsciiBtn.addEventListener("click", () => {
    const asciiContent = asciiOutput.textContent
    downloadTextFile(asciiContent, "ascii-art.txt")
  })

  copyAsciiBtn.addEventListener("click", () => {
    const asciiContent = asciiOutput.textContent
    navigator.clipboard.writeText(asciiContent).then(() => {
      // Feedback visuel : changer temporairement le texte du bouton
      const originalTitle = copyAsciiBtn.title
      copyAsciiBtn.title = "Copié !"
      setTimeout(() => {
        copyAsciiBtn.title = originalTitle
      }, 2000)
    }).catch(() => {
      alert("Erreur lors de la copie")
    })
  })

  // Fonction pour télécharger un fichier texte
  function downloadTextFile(content, filename) {
    const blob = new Blob([content], { type: "text/plain" })
    const url = URL.createObjectURL(blob)
    const link = document.createElement("a")
    link.href = url
    link.download = filename
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    URL.revokeObjectURL(url)
  }
})
