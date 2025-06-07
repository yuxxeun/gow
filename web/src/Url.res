@react.component
let make = () => {
  let (bacon, setBacon) = React.useState(() => Bacon.Loading)

  let handleBacon = () => Bacon.fetch()->then(t => setBacon(_ => t)->resolve)

  React.useEffect0(() => {
    let _ = handleBacon()
    None
  })

  <div className="p-6">
    <h1 className="text-3xl font-semibold mb-5"> {"Bacon Ipsum"->React.string} </h1>
    <Button
      onClick={_ => {
        let _ = handleBacon()
      }}>
      {"refresh data"->React.string}
    </Button>
    {bacon->Bacon.render}
  </div>
}
