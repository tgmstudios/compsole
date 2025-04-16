import { useEffect, useRef } from 'react';

export default function OpenStackNoVNC({ consoleUrl }: { consoleUrl: string }) {
  const iframeRef = useRef<HTMLIFrameElement>(null);
  
  // Process URL and force remote cursor parameter
  const secureConsoleUrl = (() => {
    const url = new URL(consoleUrl.replace(/^http:/i, 'https:'));
    url.searchParams.set('UI.cursor', 'remote');
    url.searchParams.set('UI.viewOnly', 'false');
    return url.toString();
  })();

  useEffect(() => {
    const iframe = iframeRef.current;
    if (!iframe) return;

    const handleMouseEnter = () => {
      document.documentElement.style.cursor = 'none';
    };

    const handleMouseLeave = () => {
      document.documentElement.style.cursor = 'default';
    };

    iframe.addEventListener('mouseenter', handleMouseEnter);
    iframe.addEventListener('mouseleave', handleMouseLeave);

    return () => {
      iframe.removeEventListener('mouseenter', handleMouseEnter);
      iframe.removeEventListener('mouseleave', handleMouseLeave);
    };
  }, []);

  return (
    <iframe
      ref={iframeRef}
      id="console"
      title="console"
      src={secureConsoleUrl}
      style={{
        width: '100%',
        flexGrow: '1',
        display: 'block',
        cursor: 'none'
      }}
      allowFullScreen
    />
  );
}
