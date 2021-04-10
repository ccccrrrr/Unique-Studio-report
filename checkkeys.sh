read -p "please enter one character: " C
case "$C" in
[a-z]|[A-Z])
echo "this is english character"
;;
[0-9])
echo "this is number"
;;
*)
echo "this is other token"
esac

