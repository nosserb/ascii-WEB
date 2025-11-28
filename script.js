document.addEventListener("DOMContentLoaded", () => {
  const textInput = document.getElementById("text-input")
  const generateButton = document.getElementById("generate-button")
  const asciiOutput = document.getElementById("ascii-output-text")
  const colorCodeOutput = document.getElementById("color-code-output")
  const colorOptions = document.querySelectorAll(".color-option")
  const downloadAsciiBtn = document.getElementById("download-ascii")
  const downloadColorBtn = document.getElementById("download-color")

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

      // Mettre à jour le texte de sortie
      colorCodeOutput.value = `Couleur Hex: ${currentColor}`

      // Appliquer la couleur au texte ASCII
      asciiOutput.style.color = currentColor
    })
  })

  // Initialiser la couleur et la sélection (pour le blanc par défaut)
  const defaultColorOption = document.querySelector('.color-option[data-color="#ffffff"]')
  if (defaultColorOption) {
    defaultColorOption.classList.add("selected")
  }
  colorCodeOutput.value = `Couleur Hex: ${currentColor}`
  asciiOutput.style.color = currentColor

  // --- Gestion du bouton GÉNÉRER ---
  generateButton.addEventListener("click", () => {
    const textToConvert = textInput.value

    if (textToConvert.trim() === "") {
      asciiOutput.textContent = "Veuillez entrer du texte à convertir."
      return
    }

    // Exemple de sortie simulée
    const simulatedAsciiArt = `
╔═══════════════════════╗
║   ${textToConvert.toUpperCase()}   ║
╚═══════════════════════╝
        `

    // Afficher l'ASCII Art généré
    asciiOutput.textContent = simulatedAsciiArt.trim()
    asciiOutput.style.color = currentColor
  })

  // --- Gestion du bouton Police ASCII ---
  document.getElementById("font-select-button").addEventListener("click", () => {
    alert("Sélection de police ASCII à implémenter.")
  })

  // --- Gestion des boutons de Téléchargement ---
  downloadAsciiBtn.addEventListener("click", () => {
    const asciiContent = asciiOutput.textContent
    downloadTextFile(asciiContent, "ascii-art.txt")
  })

  downloadColorBtn.addEventListener("click", () => {
    const colorContent = colorCodeOutput.value
    downloadTextFile(colorContent, "color-code.txt")
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
