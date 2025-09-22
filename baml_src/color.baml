enum Colors {
    RED
    GREEN
    YELLOW
}

enum Food {
    FRUIT
    DAIRY
    VEGETABLE
}

enum Tribe {
    Bedonkohe
    Lakota
    Sauk
    Nez_Perce
}

enum Task {
    ACTIONABLE
    NON_ACTIONABLE
}

enum Feelings {
    HAPPY
    SAD
}

enum Category {
    @@dynamic
}

function Categorize(input: string) -> Category {
    client "openai/gpt-4o"
    prompt #"
        Classify it
        {{ input }}

        {{ ctx.output_format }}
    "#
}

test leaf {
  functions [Categorize]
  args {
    input #"leaf"#
  }
}

test santa {
  functions [Categorize]
  args {
    input #"Santa Claus"#
  }
}
