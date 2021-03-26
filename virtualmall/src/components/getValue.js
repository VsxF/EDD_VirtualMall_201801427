function getValuee(node) {
    if (node == null || node == undefined) {
        return null
    } 
    if (node.kind != null) {
        node = getValuee(node.value)
    }
    if (typeof(node) === 'number') {
        return null
    }
    return node
}

export default getValuee;