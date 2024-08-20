// Settings
const PREFIX = "ext";
const SEPARATOR = ".";
const SUFFIX_RANDOM_SET = "abcdefghijklmnopqrstuvwxyz0123456789";
const SUFFIX_LENGTH = 8;

// Elements
const formInput = document.getElementById("form-input");
const inputExternalParty = document.getElementById(
  "input-external-party"
) as HTMLInputElement;
const inputDomain = document.getElementById("input-domain") as HTMLInputElement;
const inputResult = document.getElementById("input-result") as HTMLInputElement;
const btnCopyToClipboard = document.getElementById("btn-copy-to-clipboard");

formInput?.addEventListener("submit", processForm);

function processForm(event: Event) {
  // Prevent actual submit
  event.preventDefault();

  // Get input values
  const externalParty = inputExternalParty?.value;
  const domain = inputDomain?.value;

  // Generate result
  const result = generateResult({
    prefix: PREFIX,
    separator: SEPARATOR,
    externalParty,
    suffixRandomSet: SUFFIX_RANDOM_SET,
    suffixLength: SUFFIX_LENGTH,
    domain,
  });

  // Set result
  inputResult.value = result;
}

export interface IConfig {
  prefix: string;
  separator: string;
  externalParty: string;
  suffixRandomSet: string;
  suffixLength: number;
  domain: string;
}

export function generateResult(config: IConfig): string {
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

function randomStringElement(input: string): string {
  const randomBuffer = new Uint32Array(1);
  window.crypto.getRandomValues(randomBuffer);
  return input[randomBuffer[0] % (input.length - 1)];
}

btnCopyToClipboard?.addEventListener("click", copyToClipboard);

function copyToClipboard() {
  if (navigator && navigator.clipboard) {
    navigator.clipboard.writeText(inputResult.value);
    console.log("Copying text using navigator.clipboard");
  } else {
    console.log("Copying text using document.execCommand");
    inputResult.focus();
    inputResult.select();
    document.execCommand("copy");
  }
}
