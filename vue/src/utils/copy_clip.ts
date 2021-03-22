export const copyToClip = (content: string) => {
  const aux = document.createElement("textarea");
  // aux.setAttribute("value", content); 
  aux.value = content
  document.body.appendChild(aux)
  aux.select()
  document.execCommand("copy")
  document.body.removeChild(aux)
}