"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.generateResult = generateResult;
function generateResult(config) {
    // Prepare
    let parts = [];
    const cleanedParty = config.externalParty
        .trim()
        .replaceAll(" ", config.separator);
    // Add prefix if defined
    if (config.prefix != "") {
        parts.push(config.prefix);
    }
    // Add external party
    parts.push(cleanedParty);
    // Add suffix
    let suffix = "";
    for (let i = 0; i < config.suffixLength; i++) {
        suffix += randomStringElement(config.suffixRandomSet);
    }
    parts.push(suffix);
    const username = parts.join(config.separator);
    // Add domain
    if (config.domain != "") {
        return username + "@" + config.domain.trim();
    }
    return username;
}
function randomStringElement(input) {
    const randomBuffer = new Uint32Array(1);
    window.crypto.getRandomValues(randomBuffer);
    return input[randomBuffer[0] % (input.length - 1)];
}
