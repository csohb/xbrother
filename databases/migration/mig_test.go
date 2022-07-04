package main

import (
	"testing"
	"xbrother/databases/entities"
	"xbrother/databases/mysql"
)

func TestMember(t *testing.T) {
	err := mysql.ConnectMysql()
	if err != nil {
		t.Error(err)
	}

	member := []entities.TBMember{
		{
			Name:  "Mark Lee",
			Score: 3,
		}, {
			Name:  "이동혁",
			Score: 4,
		}, {
			Name:  "이제노",
			Score: 5,
		}, {
			Name:  "황인준",
			Score: 1,
		}, {
			Name:  "나은혈",
			Score: 5,
		}, {
			Name:  "철수",
			Score: 0,
		}, {
			Name:  "쌈마",
			Score: 0,
		}, {
			Name:  "공고",
			Score: 10,
		},
	}

	if err = mysql.Engine.Create(&member).Error; err != nil {
		t.Error(err)
	}
}

func TestLines(t *testing.T) {
	err := mysql.ConnectMysql()
	if err != nil {
		t.Error(err)
	}

	lines := []entities.TBLines{
		{
			Line:       "명찰 -2 웃기니까 +2 ㅋㅋ몰라 +10",
			MemberCode: 1,
		}, {
			Line:       "너 2개국어 하니 쟤는 해 새겨 들어.",
			MemberCode: 2,
		}, {
			Line:       "이쀼리들 헤쳐 모여.",
			MemberCode: 2,
		}, {
			Line:       "인순이형 참으세요! 또 정학당하면 어쩌려고 그래요.",
			MemberCode: 4,
		}, {
			Line:       "우리 학교에서 유일하게 영어 스펠링 명찰단 형.",
			MemberCode: 1,
		}, {
			Line:       "난 발기맛이 제일 좋아요. 형을 떠올리게 하니까요. 발기처럼 달콤한 핑크빛 제 마음이 느껴져요?",
			MemberCode: 2,
		}, {
			Line:       "우리 인준이 한테 마크가 있어서 다행이야.",
			MemberCode: 1,
		}, {
			Line:       "좀만 일찍 고백하지.\n왜?\n내 글 여주 모델을 너로 삼을걸",
			MemberCode: 4,
		}, {
			Line:       "영어 1등급 받겠다고.",
			MemberCode: 2,
		}, {
			Line:       "알 좀 빌려줘. 쌈마야.",
			MemberCode: 4,
		}, {
			Line:       "이동혁이 여기 없는 걸 감사해.\n그리고 오르지도 못할 자리 넘보지 마. \n떨어질라.",
			MemberCode: 3,
		}, {
			Line:       "인준아. 형은 너희 집에서.\n형 집이에요.",
			MemberCode: 4,
		}, {
			Line:       "야 황.",
			MemberCode: 2,
		}, {
			Line:       "인준아. 난 너 좋아. 헛소리에 놀아나지 않아서. 없는 권력에 기생하지 않아서. \n그러니까 널 구제해줄 사람은 너밖에 없어.",
			MemberCode: 5,
		}, {
			Line:       "난 콜라 대신 오렌지 주스잖아.",
			MemberCode: 5,
		}, {
			Line:       "이동혁은 미친 듯이 내 위로 계속 기어 올라왔다.\n처절할 정도로 맹목적이었다.\n네가 지킬 수는 있냐던 빈정을 행동으로 반박했다.",
			MemberCode: 2,
		}, {
			Line:       "요맨큼 빼고 전부 다.",
			MemberCode: 5,
		}, {
			Line:       "나는 형을 사랑해. 나는 나재민을 사랑해. 나는 이제노를 사랑해.\n나는 이동혁을.",
			MemberCode: 4,
		}, {
			Line:       "그럼 또 어쩔 건데.\n사랑은 다 소용없어. 내가 인터넷 소설만 봤어도 이건 알아.\n나중에 가면 다 헛짓이야. 쓸데없는 봉사고 적선이고 동정이라고.\n사랑이면 그 짓이 미화될 것 같아? 나중에 땅을 치고 후회해.\n아무 탈 없는 인간이나 만나서 행복하게 살아야 할 거 아냐. 캐나다 여권들고 떠나야 할 거 아냐.\n무슨 꼴 보겠다고 조센징. 그래 씨발 조센징 좆만한 땅에 비비고 살어. 형 미쳤어?",
			MemberCode: 4,
		}, {
			Line:       "동혁이네 가면 나도 동글이가 된 것 같아.",
			MemberCode: 1,
		}, {
			Line:       "형은 형이지. 형이 누구 꺼가 되고 말고가 어딨어요.",
			MemberCode: 2,
		},
	}

	if err = mysql.Engine.Create(&lines).Error; err != nil {
		t.Error(err)
	}

}
