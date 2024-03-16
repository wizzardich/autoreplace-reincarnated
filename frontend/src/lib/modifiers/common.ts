export type ModifierFunction<T> = (input: T) => T;

export type Modifier<T> = {
    enabled: boolean;
    action: ModifierFunction<T>;
}