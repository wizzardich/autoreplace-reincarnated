export class Replacement {
    from: string;
    to: string;

    constructor(from: string, to: string) {
        this.from = from;
        this.to = to;
    }
}

export class Template {
    id: string;
    name: string;
    replacements: Replacement[];

    constructor(id: string, name: string, replacements: Replacement[]) {
        this.id = id;
        this.name = name;
        this.replacements = replacements;
    }

    addReplacement(from: string, to: string) {
        this.replacements.push(new Replacement(from, to));
    }
}