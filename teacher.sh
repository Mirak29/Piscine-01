export KEY_INTERVIEW_NUMBER=$(head -n 179 ./streets/Buckingham_Place | tail -n1 | cut -d '#' -f2)
echo $KEY_INTERVIEW_NUMBER
cat interviews/interview-$KEY_INTERVIEW_NUMBER
echo $MAIN_SUSPECT