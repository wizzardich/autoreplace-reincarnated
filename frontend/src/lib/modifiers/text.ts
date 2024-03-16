export function NewLineAugmenter(input: string): string {
    return input.replaceAll('\n', '  \n');
}