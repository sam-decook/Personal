# These two programs could be combined to accept a parameter. The code is an almost exact duplicate

import sys          #for sys.exit()

def isUnique(four):
    for i in range(4):
        for j in range(i+1, 4):
            if four[i] == four[j]:
                return False
    return True


input = "cvtvbvfbvfbfqfrrtnnmqqmcqqpfplfpphchpcpcfppmwwhhscccjwjnnctcjtjgjbgjjdzjjwzwfffnhhgppjmpmgppppzbzbrzznttfqqhtqqbssrjjffgssnqnsqnsqnnmlmmztmtqmqccqscsshrsrtstrrjwrrswwpnnftfwwwzznsstbbbtzzncnbcnbccwtwctwwjgjcjtccblclgcgjgngwgbggpbbbbbgcczzfsfbbwlwwnqwnwbnwnmmbjjnfnqfqfvvgvqqvcvggpnpssptpnpjjvvpdpqpcqpcqqnhhnqncqclcfcfjcjhjdjqddhndnmmvmfmsmdsdwsdswsnsgsmsrmrjjpzzzzcppczcscsgcczqcqnnpzpbbdqqqqsqlslzljzjdjbjdjvvlhvvnbbjtbbqbdqdggzwgghhdssmttjtttmwtmtjjcpjpjptjpjmpptnntgttjpttljjwtjwwrgrvgrvrmvvrvdrdcdmmvvbpbddfbdbppdccjdjvjpjjgzjzdzfddqcqlccbvbqvqvjqjggqrqdrqrhhwggmvgmgvgrvvzwzddlzlbbdpbdbtbhthnthttjwtwffwvvrccnpcnclcgcrrgdrrgzgmzggzhhftfppnwppqhhmddzhdzzzjqjhqqcgqqmjjwppfrprlrjlrjjlhjlhjjnbnlblbzlbbfjjwfjjdcdndvvbllfmlmfmvmnnlcnncsnslsvlsvlsltslslqsqzssslqlssqrrpdpwdppsfpsprpffchhcvcpvvjpjvjmvvrjrqrzqqqlvvzzggcgrcgcmmvnmmchmmswsvwsvwswnwswdwmwzmmqfqrqtrqqwqlqlmqmsmggpzpjjlvllzplzlbzzbhhglhhwwhfwwpwgwrwlwplppfvflftltffccmcrcgrcgghgmmbrbppwvvghvhphhztzpzzwbbtpppcnpccnvcvbccvfvmffmtmbttdqdrrfnrnffsgstttbhhgmmtffhthllsbllsqsllptpwpvpdpsddsgsrrngnpggwbwwppfqfmqfmfvfzzzldltdtvdvtvtbvtgpvtnvcsdbwjrqwrfzdhwfwnmtrltdhsfzwwfpscrlsffchnfszqdddsztwzcbwdwfnljpfdsqgpnjffqfgsprggfwgvfldvzmnrpnzvfscswnzbljtvshzmcztqqwtzdstsqggqqvvlmspjjcmllcmsndwpllrvlrglphcqjgwzntwqffthrfmfgpcvjvclvdrjrpnwtrlshlnwgbbznpflsfmgvbwmnwjpnqhmdnsmwnmblqzjfzfcslhntnjhbrvvlfbffsbhrtvjltqrqszsrpmdghmhsrqlmrjdsgtsgwggdlqptqrvgmwrpcjrnmmmcdpgjzgghqmshhmddfvfrqwnwvwbvzzpmwzmvvgsqlsgvmtczwnbzrtnzjnpjjghdzsmcgvqpdhltvltbfvgsgzhvqjjnpspngqpvqwvtqrczqzwrclqbnplmvlnvcwdbzlrnwbppvdfzmblnqfpfqlngttrrcjzfjnrnmssdrrvbvhtrfzblzsjbqwbttsshmcmfhrfbfqspvcfmqbwjszmvwqnpszcfhtnqtftvtpvcjrcnlmfnvlpjzspcnmvvshpsdfhtbhgjcmjmtpmtcbwnlgtnvlwlggbnzfrrqljmjrwjtzfmgbllmmjtfhrszfrfpftmpzhzdhfwdbdnwqspjbrlglgdwtfzcrdhwgqltvthbnrqfngjgnddqlcpjflqzvbhmlzsvpvlzdvhspmchrjlzfgnlpwdnszqdqpnrchdrswfwnflslcpfmbfrplsmchbgjhhhlzjlvztpdhslrpmmcthhjwfczjtmbmsvzvsqrvzrcnflqldfnhjzlcsfhgrtqtscnbzqjglgjnqbtprrprrdtqqtlqwdqnmgbglvwtrpnpjqdsmvczbnldvmgmtpmmqbwgjbfsqbjvnjhhlrppptlrjsptpvghbgtvwvsrnggznjsbjpnwpdsqmmjjqjdwshqwzvqgpnlhdldwzmvmdnhrnhpscrbnbdtmfqwlmrlszljpjdqgjpgcdwtqzsbqsndjdslrstjwbbfnjvrjlbrphqspltjzpvsgccfrmqhfhjmlslncpgpflhjlrzqsffrtcgsvzpjmlqjznscgjrjhzfjbbchvgjgvfsrpnzdrstdtsltqldfldvfdjqdwsltcbnnvcgvnjpgdjzsbzgfvgvtmghwfmrmblscfdqlsdgcdbgsphfvvvnffcqwlrvnsrjmvvdrjdljlppsfsrmwwvmsqcgtgtfjhqbqvlsnwdwzhncgprbspjzqnbllwcjnqpfzlhscdtvlldffngwttntsbtlctdfbzjgclgbhhjtlffvtgbdsswczlgmwntrrnnbdtflplfbtqcfrhtgjnrltvzqlcvtdthldqrvjdmlntdjzcncrplgmdbqhlrrfdszrnbrllhgpzzgjnqzrjrtjjntqvdphgvjtbtqfjwvlpprrcvmpjcdmntvtqjvjrhngbtmdtflpqmlmmvjjjpdlcvntlvlgstbblwfwjrbcqjgzvjjzflfhdvgnhpnqprcvmhsrpmgfcgjpdbqpmpvlnmzdngzqcqtdsbbqbdhtgpqmqwtcctsvhwmlhpgzgpwwsqngsrwzghfhzbqcpspfhqbllbvpwgwdnghqnvjtfqtwgvnbfbpcffbwwbfqgqnftpqdtwhmrprpgdlnwrcmdndllwbvlrrjvmwlbgfwgbhtmbqhfvhlhlthbrljwnwcgsdgsbbzgnwzmjfgnfjcqhspjwctntznjptsnbpwtfzmrqwjzrmdthfwdvttrcghbjndfmtswhsmdwwwffmgsvshmlhdbwdscjnvrvqnwttlsmfbssnsscftcrpwfbbplbpmvzmttvbgpnjtqqzcfdfwqdphggwbnhfvssmdjvbvtlrcghltthqrqdbvjnbqpzjcggrbnrswhvdcrmgncsmhsmstmcbwgbfwvzqtdrpjbgljdgcncgcdslhftptnvfhgpsjqqmnsjmqhztnfzqqqhccbnqjlhqcwwbfzgpmbvjszswjlhcqshtqdmwmhpgqgzbwbtchrdwgnvfqcgcmngtddpzfgjrrssdmcnbtpbhfvfbdgzmcbjqtdqrnrdchtrcfpmtnvfczdlnjplwccfmzjwwvdsrcjznrcbczdqwshjgvhjpgdgjvtvmdzncrqgffzhcbwstvltqbrzfbnrrvlntjmzctvnqhdcgvvmmbfplzpwgsdzdtvnwnshnwbwqplbctntslczwhblsnzdcsgltvlfvbtvthvlsmzmfgzhpvrvczdhmcpmgqzbdhbcdbstrrtzgrbhhwwghzlgcdtlcjtjdcwdtqpvpwzrwmhczfsdmwjlgtmzvdbwzdnsztvqzhjpfrqzllgvrsbmhvrlzwtfmpzchftrzlmndpwqgcmhdzmjrjvpbnrcmjvzbzlqngmrzcbdwfcwmlnlqsdgbvmvqcsjtwgvtpmqdtdqfrshmlgnjqdwsrpjcfnfpjzlhslsvppmgprqlpcfztcngmqgrvqvpcszzgtbpcrghmgnvgplctgdljhwjgbzmjrdsccdtrwjrrbttzmpvmzlmqdfwdccfzqztdjmcgrcdwghgmmntdwrclshlgnsmhmltwcczncndtglnqvcvbwlvcgqtzpnqnlllchwjwrlhcdlbjlzvhvtjsttqhvhfqbsqjpfcmpsfshgrrwgprjtfnhsjqqntrjfgjjmjbqwdjlcjjdljpppgvqvgrrwvclbmncmqglhbtjcfrbjmhqgmlpzgljntqvrcggmrgtbrftvhrnjbvzstlnzrqcgncnjhdrnqcnwcrstvdsgnwzjzmqsdtwzzjmfgzpslhqclhtjnrlwwfrqgmhsttfdprhgmphzfdfqntwztltnnfqwgdnnbrlgzmtwqsvhbjjqtcgmghnbchdfgfrdtjnvlgtqdgrfwmgnmpslqmqzdgqztnppbc"

four = ['c','v','t','v']

for x in range(4, len(input)):
    if len(four) >= 4:
        four.pop(0)
    four.append(input[x])

    # Check if all are unique
    if isUnique(four):
        print("Characters processed: " +str(x+1))
        sys.exit()