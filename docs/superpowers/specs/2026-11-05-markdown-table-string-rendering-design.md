# Markdown Table String Rendering Design

## Goal

Keep buffer handling at the `MarkdownTable.WriteTo()` boundary while allowing callers to consume rendered table lines directly.

## Design

`renderRow()` and `renderRowDivider()` will each construct and return one Markdown line as a string. These strings will not include line terminators.

`ToStrings()` will render the table as an ordered slice containing:

1. The header row.
2. The header divider.
3. Each data row.

`WriteTo()` will call `ToStrings()`, write each returned string to the supplied builder, and append a newline after every line. This preserves the existing `WriteTo()` output while keeping line-ending policy out of `ToStrings()`.

## Compatibility

Column width and alignment behaviour will remain unchanged. Existing golden tests for `WriteTo()` will continue to verify the complete rendered output.

## Testing

A focused unit test will verify that `ToStrings()` returns the expected header, divider, and data rows as separate strings without trailing newlines. Existing golden tests will verify that `WriteTo()` still produces byte-for-byte compatible output.
