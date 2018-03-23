try
  call writefile(split(matchstr($VIMUAR_TEXT, $VIMUAR_PATTERN), "\n"), $VIMUAR_FILE)
finally
  qall!
endtry
