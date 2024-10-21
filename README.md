# OrdoBudget

Automate budget in google sheets

## Usage

- make excrptGrpData.json for matching excrpts
- make a parent grp for each groupe of expenses, i.e fixed expenses, Car expenses, insurance, Loans, variable expenses ...
-

## Entry 

```
type Entry struct {
    // Used to make lookup in excerptMappings array
    ind int

    // Name of the entry
    name string

    // Matches for this entry
    mappings []string

    // Defines the type of this excerpt
    GroupName string

    // Determines if the initial group total value should be read from the sheet or start at 0
    // Default is false
    fixedExpense bool
}
```

- fixedExpense
    - Is used for fixed expenses which rarely change

- Group: Ignored
    - Matches are not ignored.


## ExcrptData template

- Use excrptGrpDataTemplate.json as template for matching excpts


## Links

- https://docs.google.com/spreadsheets/d/1Dg3qfLZd3S2ISqYLA7Av-D3njmiWPlcq-tQAodhgeAc/edit?gid=1685114351#gid=1685114351

## Gallerie

