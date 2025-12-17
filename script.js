// --- Gestion du Mode Jour/Nuit et Génération ASCII ---
document.addEventListener("DOMContentLoaded", () => {
  // ===== GESTION DU THÈME JOUR/NUIT =====
  const themeCheckbox = document.getElementById("theme-checkbox")
  
  // Vérifier s'il y a une préférence sauvegardée
  const savedTheme = localStorage.getItem("theme") || "dark"
  
  // Appliquer le thème sauvegardé
  if (savedTheme === "light") {
    document.body.classList.add("light-mode")
    themeCheckbox.checked = true
  } else {
    document.body.classList.remove("light-mode")
    themeCheckbox.checked = false
  }
  
  // Écouter les changements du bouton
  themeCheckbox.addEventListener("change", () => {
    if (themeCheckbox.checked) {
      document.body.classList.add("light-mode")
      localStorage.setItem("theme", "light")
    } else {
      document.body.classList.remove("light-mode")
      localStorage.setItem("theme", "dark")
    }
  })

  // ===== GESTION DE LA GÉNÉRATION ASCII ART =====
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
      // Nettoyer le texte : remplacer \r\n par \n, puis \r par \n
      const cleanedText = textToConvert.replace(/\r\n/g, "\n").replace(/\r/g, "\n")

      console.log("Envoi au serveur:", { text: cleanedText, banner: selectedBanner })

      const response = await fetch("/ascii-art", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          text: cleanedText,
          banner: selectedBanner
        })
      })

      if (!response.ok) {
        const errorText = await response.text()
        console.error("Erreur serveur:", errorText)
        throw new Error(`Erreur HTTP: ${response.status} - ${errorText}`)
      }

      const asciiArt = await response.text()
      console.log("Réponse du serveur reçue")
      
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
    
    // Essayer d'abord avec l'API Clipboard moderne
    if (navigator.clipboard && navigator.clipboard.writeText) {
      navigator.clipboard.writeText(asciiContent).then(() => {
        // Feedback visuel : changer temporairement le texte du bouton
        const originalTitle = copyAsciiBtn.title
        copyAsciiBtn.title = "Copié !"
        setTimeout(() => {
          copyAsciiBtn.title = originalTitle
        }, 2000)
      }).catch((err) => {
        console.error("Erreur clipboard:", err)
        fallbackCopy(asciiContent)
      })
    } else {
      // Fallback pour les anciens navigateurs
      fallbackCopy(asciiContent)
    }
  })

  // Fallback pour la copie (compatible avec Firefox et autres)
  function fallbackCopy(text) {
    const textarea = document.createElement("textarea")
    textarea.value = text
    textarea.style.position = "fixed"
    textarea.style.opacity = "0"
    document.body.appendChild(textarea)
    textarea.select()
    try {
      document.execCommand("copy")
      const originalTitle = copyAsciiBtn.title
      copyAsciiBtn.title = "Copié !"
      setTimeout(() => {
        copyAsciiBtn.title = originalTitle
      }, 2000)
    } catch (err) {
      console.error("Erreur copie:", err)
      alert("Erreur lors de la copie")
    }
    document.body.removeChild(textarea)
  }

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
