import { generateResult, IConfig } from "../main";

describe("testing generateResult", () => {
  const testCases: [string, { given: IConfig; expected: RegExp }][] = [
    [
      "all provided",
      {
        given: {
          prefix: "ext",
          separator: ".",
          externalParty: "test",
          suffixRandomSet: "abc",
          suffixLength: 8,
          domain: "example.com",
        },
        expected: /^ext\.test\.[a-c]{8}@example\.com$/,
      },
    ],
    [
      "minimal provided",
      {
        given: {
          prefix: "",
          separator: "",
          externalParty: "test",
          suffixRandomSet: "123",
          suffixLength: 8,
          domain: "",
        },
        expected: /^test[1-3]{8}$/,
      },
    ],
    [
      "no prefix",
      {
        given: {
          prefix: "",
          separator: ".",
          externalParty: "test",
          suffixRandomSet: "123",
          suffixLength: 8,
          domain: "example.com",
        },
        expected: /^test\.[1-3]{8}@example\.com$/,
      },
    ],
    [
      "no separator",
      {
        given: {
          prefix: "ext",
          separator: "",
          externalParty: "test",
          suffixRandomSet: "123",
          suffixLength: 8,
          domain: "example.com",
        },
        expected: /^exttest[1-3]{8}@example\.com$/,
      },
    ],
    [
      "no domain",
      {
        given: {
          prefix: "ext",
          separator: ".",
          externalParty: "test",
          suffixRandomSet: "123",
          suffixLength: 8,
          domain: "",
        },
        expected: /^ext.test.[1-3]{8}$/,
      },
    ],
    [
      "inputs with whitespace",
      {
        given: {
          prefix: "ext",
          separator: "_",
          externalParty: " \t\nt e s t \t\n",
          suffixRandomSet: "123",
          suffixLength: 8,
          domain: " \t\nexample.com \t\n",
        },
        expected: /^ext_t_e_s_t_[1-3]{8}@example\.com$/,
      },
    ],
  ];

  it.each(testCases)("%s", (_, testCase) => {
    expect(generateResult(testCase.given)).toMatch(testCase.expected);
  });
});
