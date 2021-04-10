if [ $1 -ge 85 ] && [ $1 -le 100 ] ; then
echo "$1 is excellent!"
elif [ $1 -ge 75 ] ; then
echo "$1 is good!"
elif [ $1 -ge 60 ] ; then
echo "$1 is ok"
elif [ $1 -ge 0 ] ; then
echo "fail..."
fi
