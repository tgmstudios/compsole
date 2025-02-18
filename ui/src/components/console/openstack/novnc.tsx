export default function OpenStackNoVNC({ consoleUrl }: { consoleUrl: string }) {
  const secureConsoleUrl = consoleUrl.replace(/^http:\/\//i, "https://");

  return (
    <iframe
      id="console"
      title="console"
      src={secureConsoleUrl}
      style={{
        width: '100%',
        flexGrow: '1',
        display: 'block',
      }}
    />
  )
}
